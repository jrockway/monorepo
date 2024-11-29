package main

import (
	"archive/tar"
	"context"
	"flag"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/jrockway/monorepo/internal/errors"
	"github.com/jrockway/monorepo/internal/links"
	"github.com/klauspost/compress/zstd"
)

var (
	out  = flag.String("out", "/tmp/out.tar.zst", "location to put the built site archive")
	cd   = flag.String("cd", "", "directory to run hugo inside of")
	hugo = "hugo" // replaced by the linker
)

func run(ctx context.Context, cd, out string) (retErr error) {
	outdir, err := os.MkdirTemp("", "hugo")
	if err != nil {
		return errors.Wrap(err, "make output dir")
	}
	defer errors.Invoke1(&retErr, os.RemoveAll, outdir, "cleanup output dir")

	h, err := runfiles.Rlocation(hugo)
	if err != nil {
		return errors.Wrapf(err, "lookup rlocation of hugo (%v)", hugo)
	}
	h, err = filepath.Abs(h)
	if err != nil {
		return errors.Wrapf(err, "lookup absolute path of hugo (%v)", h)
	}

	indir, err := os.MkdirTemp("", "input")
	if err != nil {
		return errors.Wrap(err, "make input dir")
	}
	defer errors.Invoke1(&retErr, os.RemoveAll, indir, "cleanup input dir")
	if err := links.ResolveDirectory(cd, indir); err != nil {
		return errors.Wrapf(err, "resolve links from input dir %v into %v", cd, indir)
	}

	cmd := exec.CommandContext(ctx, h, "--noBuildLock", "--destination", outdir, "--environment", "production")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Join(indir, cd)
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "run hugo")
	}

	fw, err := os.Create(out)
	if err != nil {
		return errors.Wrapf(err, "create output file %v", out)
	}
	defer errors.Close(&retErr, fw, "close output file %v", out)

	zw, err := zstd.NewWriter(fw, zstd.WithEncoderLevel(zstd.SpeedBetterCompression))
	if err != nil {
		return errors.Wrapf(err, "new zstd writer")
	}
	defer errors.Close(&retErr, zw, "close zstd writer")
	tw := tar.NewWriter(zw)
	defer errors.Close(&retErr, tw, "close tar writer")

	if err := filepath.WalkDir(outdir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrapf(err, "WalkDir called with error on path %v", path)
		}
		if path == outdir {
			return nil
		}
		rel, err := filepath.Rel(outdir, path)
		if err != nil {
			return errors.Wrapf(err, "calculate relative path of %v", path)
		}
		rel = filepath.Join("srv", rel)
		info, err := d.Info()
		if err != nil {
			return errors.Wrapf(err, "get info for %v", path)
		}
		if d.IsDir() {
			if err := tw.WriteHeader(&tar.Header{
				Typeflag: tar.TypeDir,
				Name:     rel,
				Mode:     int64(info.Mode()),
				Uid:      65534,
				Gid:      65534,
				ModTime:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			}); err != nil {
				return errors.Wrapf(err, "add directory %v", rel)
			}
			return nil
		}
		if err := tw.WriteHeader(&tar.Header{
			Typeflag: tar.TypeReg,
			Name:     rel,
			Mode:     int64(info.Mode()),
			Size:     info.Size(),
			Uid:      65534,
			Gid:      65534,
			ModTime:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		}); err != nil {
			return errors.Wrapf(err, "add file %v", rel)
		}
		fh, err := os.Open(path)
		if err != nil {
			return errors.Wrapf(err, "open %v", path)
		}
		if _, err := io.Copy(tw, fh); err != nil {
			return errors.Wrapf(err, "copy %v into archive (%v)", path, rel)
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "walk output tree")
	}
	if err := tw.Close(); err != nil {
		return errors.Wrap(err, "close tar writer")
	}
	if err := zw.Close(); err != nil {
		return errors.Wrap(err, "close zstd writer")
	}
	return nil
}

func main() {
	flag.Parse()
	ctx, c := signal.NotifyContext(context.Background(), os.Interrupt)
	defer c()
	if err := run(ctx, *cd, *out); err != nil {
		log.Fatalf("run: %v", err)
	}
}
