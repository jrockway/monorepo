// Package tarutil contains utilities for working with ye olde tape archives.
package tarutil

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/jrockway/monorepo/internal/errors"
	"github.com/jrockway/monorepo/internal/log"
	"github.com/zeebo/blake3"
	"go.uber.org/zap"
)

// Unpack unpacks a tar archive to os.TempDir, preferring a cached copy.
func Unpack(ctx context.Context, name, rlocation string) (_ string, retErr error) {
	loc, err := runfiles.Rlocation(rlocation)
	if err != nil {
		return "", errors.Wrapf(err, "resolve archive rlocation (%v)", rlocation)
	}
	h := blake3.New()
	r, err := os.Open(loc)
	if err != nil {
		return "", errors.Wrapf(err, "open input archive for hash")
	}
	defer errors.Close(&retErr, r, "close input %v", loc)

	if _, err := io.Copy(h, r); err != nil {
		return "", errors.Wrapf(err, "hash input archive")
	}
	hash := h.Sum(nil)
	dst := filepath.Join(os.TempDir(), fmt.Sprintf("com-github-jrockway-monorepo-%v-%x", name, hash))
	if st, err := os.Stat(dst); err == nil && st.IsDir() {
		log.Debug(ctx, "using cached copy of unpacked archive", zap.String("name", name), zap.String("path", dst))
		return dst, nil
	}

	log.Debug(ctx, "no cached copy of unpacked archive; creating", zap.String("name", name), zap.String("path", dst))
	tmp, err := os.MkdirTemp("", "com-github-jrockway-monorepo-unpack-*")
	if err != nil {
		return "", errors.Wrap(err, "create tmpdir for archive unpack")
	}
	var outputOK bool
	defer func() {
		if outputOK {
			return
		}
		if err := os.RemoveAll(tmp); err != nil {
			errors.JoinInto(&retErr, errors.Wrapf(err, "remove invalid tmpdir %v", tmp))
		}
		log.Debug(ctx, "cleaned up invalid output", zap.String("path", tmp))
	}()
	if _, err := r.Seek(0, 0); err != nil {
		return "", errors.Wrap(err, "seek input to beginning for unpack")
	}
	tr := tar.NewReader(r)
	for {
		h, err := tr.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", errors.Wrap(err, "TarReader.Next")
		}
		rel := filepath.FromSlash(h.Name)
		abs := filepath.Join(tmp, rel)
		switch h.Typeflag {
		case tar.TypeReg:
			w, err := os.OpenFile(abs, os.O_RDWR|os.O_CREATE|os.O_TRUNC, h.FileInfo().Mode().Perm())
			if err != nil {
				return "", errors.Wrapf(err, "open output file %v", abs)
			}
			n, err := io.Copy(w, tr)
			if closeErr := w.Close(); closeErr != nil {
				errors.JoinInto(&err, errors.Wrap(closeErr, "close output"))
			}
			if err != nil {
				return "", errors.Wrapf(err, "write output file %v", abs)
			}
			if got, want := n, h.Size; got != want {
				return "", errors.Errorf("output size mismatch on %v; got %v want %v", abs, got, want)
			}
			mtime := h.ModTime
			if err := os.Chtimes(abs, mtime, mtime); err != nil {
				return "", errors.Wrapf(err, "adjust mtime on %v to %v", abs, mtime.Format(time.RFC3339Nano))
			}
		case tar.TypeDir:
			if err := os.MkdirAll(abs, 0o755); err != nil {
				return "", errors.Wrapf(err, "mkdir %v", abs)
			}
		case tar.TypeSymlink:
			src := h.Linkname
			if err := os.Symlink(src, abs); err != nil {
				return "", errors.Wrapf(err, "link %v -> %v", abs, src)
			}
		case tar.TypeLink:
			log.Debug(ctx, "skipping hard link", zap.String("file", rel), zap.String("linkname", h.Linkname))
		default:
			return "", errors.Errorf("unsupported type for file %v: 0x%x", rel, h.Typeflag)
		}
	}
	if err := os.Rename(tmp, dst); err != nil {
		return "", errors.Wrapf(err, "rename %v to %v", tmp, dst)
	}
	outputOK = true
	return dst, nil
}
