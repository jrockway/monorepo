// Package testpostgres runs postgres and utilities for tests.
package testpostgres

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jrockway/monorepo/internal/errors"
	"github.com/jrockway/monorepo/internal/log"
	"github.com/jrockway/monorepo/internal/pctx"
	"github.com/jrockway/monorepo/internal/tarutil"
	"go.uber.org/zap"
)

var (
	postgresArchiveRlocation string
	psqlBinaryRlocation      string // TODO: macos support
)

// Unpack unpacks postgres to os.TempDir, returning the path to the root.
func Unpack(ctx context.Context, name, rlocation string) (string, error) {
	root, err := tarutil.Unpack(ctx, "postgres", postgresArchiveRlocation)
	if err != nil {
		return "", errors.Wrapf(err, "unpack tar from %v", postgresArchiveRlocation)
	}
	return root, nil
}

// PsqlCmd returns a command that runs `psql`.
func PsqlCmd(ctx context.Context) (*exec.Cmd, error) {
	root, err := Unpack(ctx, "postgres", postgresArchiveRlocation)
	if err != nil {
		return nil, errors.Wrap(err, "unpack postgres")
	}
	cmd := exec.CommandContext(ctx, filepath.Join(root, "usr", "lib", "postgresql", "17", "bin", "psql"))
	return cmd, nil
}

func RunPostgres(ctx context.Context) (cfg *pgx.ConnConfig, retErr error) {
	// Make tmpdir with data/config.
	dir, err := os.MkdirTemp("", "com-github-jrockway-monorepo-postgres-data-*")
	if err != nil {
		return nil, errors.Wrap(err, "make postgres data directory")
	}
	cleanup := func() {
		log.Debug(ctx, "cleaning up postgres tmpdir", zap.String("path", dir))
		if err := os.RemoveAll(dir); err != nil {
			errors.JoinInto(&retErr, errors.Wrap(err, "cleanup postgres data directory"))
		}
	}
	doCleanup := true
	defer func() {
		if doCleanup {
			cleanup()
		}
	}()

	// Unpack the postgres binaries archive.
	root, err := Unpack(ctx, "postgres", postgresArchiveRlocation)
	if err != nil {
		return nil, errors.Wrap(err, "unpack postgres")
	}

	// Run initdb.
	init := exec.CommandContext(ctx, filepath.Join(root, "usr", "lib", "postgresql", "17", "bin", "initdb"),
		"-D", dir,
		"--no-instructions",
		"-A", "reject",
		"-c", "listen_addresses=",
		"--auth-local=trust",
		"--no-sync",
		"-c", "fsync=off", // huge performance drain during parallel tests
		"-c", "unix_socket_directories="+dir,
	)
	init.Stdout = log.NewWriterAt(pctx.Child(ctx, "initdb.stdout"), log.DebugLevel)
	init.Stderr = log.NewWriterAt(pctx.Child(ctx, "initdb.stderr"), log.DebugLevel)
	log.Debug(ctx, "initializing database")
	if err := init.Run(); err != nil {
		return nil, errors.Wrap(err, "init postgres database")
	}

	// Start postgres.
	log.Debug(ctx, "starting database server")
	serve := exec.CommandContext(ctx, filepath.Join(root, "usr", "lib", "postgresql", "17", "bin", "postgres"), "-D", dir)
	serve.Stdout = log.NewWriterAt(pctx.Child(ctx, "stdout"), log.DebugLevel)
	serve.Stderr = log.NewWriterAt(pctx.Child(ctx, "stderr"), log.DebugLevel)
	if err := serve.Start(); err != nil {
		return nil, errors.Wrap(err, "start postgres")
	}

	// Wait for the server to exit, in the background.
	doCleanup = false
	exitCh := make(chan error)
	go func() {
		if err := serve.Wait(); err != nil {
			cleanup()
			exitCh <- err
		}
		close(exitCh)

	}()

	// Wait for the server to accept connections (or fail to start up).
	cfg, err = pgx.ParseConfig("database=postgres host=" + dir)
	if err != nil {
		return nil, errors.Wrap(err, "hard-coded config appears invalid")
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
		return nil, errors.Wrap(err, "server did not start ok")
	}
	log.Debug(ctx, "database started ok")
	return cfg, nil
}
