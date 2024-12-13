// Package testpostgres runs postgres and utilities for tests.
package testpostgres

import (
	"context"
	"os/exec"
	"path/filepath"

	"github.com/jrockway/monorepo/internal/errors"
	"github.com/jrockway/monorepo/internal/tarutil"
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
