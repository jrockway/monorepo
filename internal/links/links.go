// Package links contains utilities for dealing with filesystem links.
package links

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/jrockway/monorepo/internal/errors"
)

// ResolveDirectory copies dir to out, replacing any symlinks in dir with the actual content of the
// file.
func ResolveDirectory(dir string, out string) error {
	if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) (retErr error) {
		if err != nil {
			return errors.Wrapf(err, "%v called with error", path)
		}
		newpath := filepath.Join(out, path)
		if d.IsDir() {
			if err := os.MkdirAll(newpath, 0o755); err != nil {
				return errors.Wrap(err, "create directory in output")
			}
			return nil
		}
		src, err := os.OpenFile(path, os.O_RDONLY, 0)
		if err != nil {
			return errors.Wrap(err, "open source")
		}
		defer errors.Close(&retErr, src, "close src")
		dst, err := os.OpenFile(newpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
		if err != nil {
			return errors.Wrap(err, "open destination")
		}
		defer errors.Close(&retErr, dst, "close dst")
		if _, err := io.Copy(dst, src); err != nil {
			return errors.Wrapf(err, "copy %v to %v", dir, path)
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "walk")
	}
	return nil
}
