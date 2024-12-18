// Package testpostgres runs postgres and utilities for tests.
package testpostgres

import (
	"context"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/jackc/pgx/v4"
	"github.com/jrockway/monorepo/internal/errors"
	"github.com/jrockway/monorepo/internal/log"
	"github.com/jrockway/monorepo/internal/pctx"
	"github.com/jrockway/monorepo/internal/tarutil"
	"go.uber.org/zap"
)

var (
	postgresArchiveRlocation string
	patchelfRlocation        string
	psqlBinaryRlocation      string // TODO: macos support
)

func postgresBin(root string, binary string) string {
	if binary == "" {
		return filepath.Join(root, "usr", "lib", "postgresql", "17", "bin")
	}
	return filepath.Join(root, "usr", "lib", "postgresql", "17", "bin", binary)
}

func postgresLib(root string) string {
	// TODO: macos
	os := "x86_64-linux-gnu"
	if runtime.GOARCH == "arm64" {
		os = "aarch64-linux-gnu"
	}
	return "LD_LIBRARY_PATH=" + strings.Join([]string{
		filepath.Join(root, "lib", os),
		filepath.Join(root, "usr", "lib", os),
	}, ":")
}

func localLdSo(root string) string {
	switch runtime.GOARCH {
	case "amd64":
		return filepath.Join(root, "/lib64/ld-linux-x86-64.so.2")
	case "arm64":
		return filepath.Join(root, "/lib/ld-linux-aarch64.so.1")
	}
	return ""
}

var patchelf = ""

func patch(ctx context.Context, interp, path string) error {
	if patchelf == "" {
		var err error
		patchelf, err = runfiles.Rlocation(patchelfRlocation)
		if err != nil {
			return errors.Wrap(err, "find patchelf binary")
		}
	}
	cmd := exec.CommandContext(ctx, patchelf, "--set-interpreter", interp, "--no-default-lib", path)
	cmd.Stdout = log.NewWriterAt(ctx, log.DebugLevel)
	cmd.Stderr = log.NewWriterAt(ctx, log.DebugLevel)
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "run patchelf")
	}
	return nil
}

// Unpack unpacks postgres to os.TempDir, returning the path to the root.
func Unpack(ctx context.Context, name, rlocation string) (string, error) {
	root, err := tarutil.Unpack(ctx, "postgres", postgresArchiveRlocation)
	if err != nil {
		return "", errors.Wrapf(err, "unpack tar from %v", postgresArchiveRlocation)
	}
	patched := filepath.Join(root, ".binaries-patched")
	if _, err := os.Stat(patched); err == nil {
		log.Debug(ctx, "skipping interpreter patch; binaries already patched")
		return root, nil
	}
	interp := localLdSo(root)
	// The binaries haven't been patched yet.  Patch them.
	bin := postgresBin(root, "")
	if err := filepath.Walk(bin, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "WalkFn called with error")
		}
		if interp == "" {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if err := patch(ctx, interp, path); err != nil {
			return errors.Wrapf(err, "patch %v", path)
		}
		return nil
	}); err != nil {
		return "", errors.Wrap(err, "patch binaries")
	}
	w, err := os.Create(patched)
	if err != nil {
		return "", errors.Wrap(err, "create .binaries-patched note")
	}
	if err := w.Close(); err != nil {
		return "", errors.Wrap(err, "close .binaries-patched note")
	}
	return root, nil
}

// PsqlCmd returns a command that runs `psql`.
func PsqlCmd(ctx context.Context) (*exec.Cmd, error) {
	root, err := Unpack(ctx, "postgres", postgresArchiveRlocation)
	if err != nil {
		return nil, errors.Wrap(err, "unpack postgres")
	}
	cmd := exec.CommandContext(ctx, postgresBin(root, "psql"))
	cmd.Env = []string{postgresLib(root)}
	return cmd, nil
}

const debianNobody = 65534

func RunPostgres(ctx context.Context) (cfg *pgx.ConnConfig, awaitCleanup func(), retErr error) {
	// Unpack the postgres binaries archive.
	root, err := Unpack(ctx, "postgres", postgresArchiveRlocation)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unpack postgres")
	}
	if err := patch(ctx, localLdSo(root), filepath.Join(root, "bin", "sh")); err != nil {
		return nil, nil, errors.Wrap(err, "fix ./bin/sh interpreter")
	}

	// Make tmpdir with data/config.
	dir, err := os.MkdirTemp("", "com-github-jrockway-monorepo-postgres-data-*")
	if err != nil {
		return nil, nil, errors.Wrap(err, "make postgres data directory")
	}
	cleanup := func() {
		log.Debug(ctx, "cleaning up postgres tmpdir", zap.String("path", dir))
		if err := os.RemoveAll(dir); err != nil {
			log.Info(ctx, "problem cleaning up postgres data directory", zap.Error(err))
		}
	}
	doCleanup := true
	defer func() {
		if doCleanup {
			cleanup()
		}
	}()

	// CI runs as root, which postgres hates, so we drop privileges to make it start up cleanly.
	// It also wants to run /bin/sh, so we just replace /bin/sh with the copy we bundle.  This
	// is a little rude.  (This is working around the fact that the vendored ld-linux.so.2
	// requires a newer libc than exists on CI machines.  Not sure how to override library
	// loading for ld itself.)
	attrs := &syscall.SysProcAttr{}
	amRoot := os.Getuid() == 0
	if amRoot {
		// Change ownership of the data directory to the user we're going to switch to.
		if err := os.Chown(dir, debianNobody, debianNobody); err != nil {
			return nil, nil, errors.Wrap(err, "chown(nobody, nobody) data tmpdir")
		}

		// Move our shell in.
		if _, err := os.Stat("/bin/sh.orig"); err != nil && os.IsNotExist(err) {
			if err := os.Rename("/bin/sh", "/bin/sh.orig"); err != nil {
				return nil, nil, errors.Wrap(err, "back up /bin/sh")
			}
			if err := os.Symlink(filepath.Join(root, "bin", "sh"), "/bin/sh"); err != nil {
				return nil, nil, errors.Wrap(err, "symlink in /bin/sh")
			}
		}

		// Setup for chroot/setuid.
		attrs = &syscall.SysProcAttr{
			Credential: &syscall.Credential{
				Uid:         debianNobody,
				Gid:         debianNobody,
				NoSetGroups: true,
			},
		}
	}

	// Run initdb.
	init := exec.CommandContext(ctx, postgresBin(root, "initdb"),
		"-D", dir,
		"--no-instructions",
		"-A", "reject",
		"-c", "listen_addresses=",
		"--auth-local=trust",
		"--no-sync",
		"-c", "fsync=off", // huge performance drain during parallel tests
		"-c", "unix_socket_directories="+dir,
		"-U", "postgres",
	)
	init.Stdout = log.NewWriterAt(pctx.Child(ctx, "initdb.stdout"), log.DebugLevel)
	init.Stderr = log.NewWriterAt(pctx.Child(ctx, "initdb.stderr"), log.DebugLevel)
	init.Env = []string{postgresLib(root), "PATH=" + filepath.Join(root, "bin")}
	init.SysProcAttr = attrs // postgres is very mad about being root, but CI is root!
	log.Debug(ctx, "initializing database")
	if err := init.Run(); err != nil {
		return nil, nil, errors.Wrap(err, "init postgres database")
	}

	// Start postgres.
	log.Debug(ctx, "starting database server")
	serve := exec.CommandContext(ctx, postgresBin(root, "postgres"), "-D", dir)
	serve.Stdout = log.NewWriterAt(pctx.Child(ctx, "stdout"), log.DebugLevel)
	serve.Stderr = log.NewWriterAt(pctx.Child(ctx, "stderr"), log.DebugLevel)
	serve.Env = []string{postgresLib(root)}
	serve.SysProcAttr = attrs
	if err := serve.Start(); err != nil {
		return nil, nil, errors.Wrap(err, "start postgres")
	}

	// Wait for the server to exit, in the background.
	doCleanup = false
	exitCh := make(chan error)
	cleanupCh := make(chan struct{})
	go func() {
		err := serve.Wait()
		cleanup()
		close(cleanupCh)
		if err != nil {
			select {
			case exitCh <- err:
			case <-ctx.Done():
			}
		}
		close(exitCh)

	}()

	// Wait for the server to accept connections (or fail to start up).
	cfg, err = pgx.ParseConfig("database=postgres user=postgres host=" + dir)
	if err != nil {
		return nil, nil, errors.Wrap(err, "hard-coded config appears invalid")
	}
	pingCh := make(chan error)
	go func(rctx context.Context) {
		defer close(pingCh)
		var startupErr error
		for i := 0; i < 30; i++ {
			ctx, cancel := context.WithTimeout(rctx, time.Second)
			select {
			case <-ctx.Done():
				pingCh <- context.Cause(ctx)
				cancel()
				return
			case err := <-exitCh:
				pingCh <- err
				cancel()
				return
			case <-time.After(100 * time.Millisecond):
				startupErr = nil
			}
			c, err := pgx.ConnectConfig(ctx, cfg)
			if err != nil {
				startupErr = errors.Wrap(err, "connect")
				cancel()
				continue
			}
			if err := c.Ping(ctx); err != nil {
				startupErr = errors.Wrap(err, "ping")
				cancel()
				continue
			}
			if err := c.Close(ctx); err != nil {
				cancel()
				continue
			}
			cancel()
			return
		}
		if startupErr != nil {
			pingCh <- errors.Wrap(err, "start postgres: after 30 attempts")
		} else {
			pingCh <- errors.New("server failed to accept connections after 3s")
		}

	}(ctx)

	// Report the result of starting up.
	if err := <-pingCh; err != nil {
		if killErr := serve.Process.Kill(); err != nil {
			errors.JoinInto(&err, errors.Wrap(killErr, "killing postgres"))
		}
		return nil, nil, errors.Wrap(err, "server did not start ok")
	}
	log.Debug(ctx, "database started ok")
	return cfg, func() {
		if err := serve.Process.Kill(); err != nil {
			if !errors.Is(err, os.ErrProcessDone) {
				log.Debug(ctx, "problem killing postgres", zap.Error(err))
			}
		}
		<-cleanupCh
	}, nil
}
