package store

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/bazelbuild/rules_go/go/runfiles"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/tern/migrate"
)

const (
	versionTable  = "public.schema_version"
	migrationPath = "migrations"
)

var (
	migrationsLocation string
)

func findMigrations() (string, error) {
	migrations := strings.Split(migrationsLocation, " ")
	if len(migrations) == 0 {
		return "", errors.New("no migrationsLocation set in x_defs; unable to locate migrations")
	}

	location, err := runfiles.Rlocation(migrations[0])
	if err != nil {
		return "", fmt.Errorf("no runfile for first migration %v: %w", migrationsLocation, err)
	}
	return path.Dir(location), nil
}

func (c *Connection) MigrateDB(ctx context.Context) error {
	conn, err := c.db.DB.Conn(ctx)
	if err != nil {
		return fmt.Errorf("get raw connection: %w", err)
	}
	err = conn.Raw(func(driverConn interface{}) error {
		conn := driverConn.(*stdlib.Conn).Conn()
		m, err := migrate.NewMigrator(ctx, conn, versionTable)
		if err != nil {
			return fmt.Errorf("new migrator: %w", err)
		}
		path, err := findMigrations()
		if err != nil {
			return fmt.Errorf("find migrations: %w", err)
		}
		if err := m.LoadMigrations(path); err != nil {
			return fmt.Errorf("load migrations from runfiles: %w", err)
		}
		if err := m.Migrate(ctx); err != nil {
			return fmt.Errorf("migrate: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}
	return nil
}
