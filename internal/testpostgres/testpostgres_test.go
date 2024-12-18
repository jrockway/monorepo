package testpostgres

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v4"
	"github.com/jrockway/monorepo/internal/pctx"
)

func TestDatabase(t *testing.T) {
	ctx := pctx.TestContext(t)
	cfg, cleanup, err := RunPostgres(ctx)
	if err != nil {
		t.Fatalf("start postgres: %v", err)
	}
	t.Cleanup(cleanup)
	c, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	if _, err := c.Exec(ctx, "create database test"); err != nil {
		t.Fatalf("create database test: %v", err)
	}
	if err := c.Close(ctx); err != nil {
		t.Fatalf("close connection to 'postgres': %v", err)
	}
	cfg.Database = "test"
	c, err = pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		t.Fatalf("connect to 'test': %v", err)
	}
	if _, err := c.Exec(ctx, "create table foo(id bigserial primary key not null, k text, v text); insert into foo(k, v) values ('key', 'value')"); err != nil {
		t.Fatalf("create table foo: %v", err)
	}

	type result struct {
		ID   int64
		K, V string
	}
	var got result
	row := c.QueryRow(ctx, "select id, k, v from foo")
	if err := row.Scan(&got.ID, &got.K, &got.V); err != nil {
		t.Fatalf("select from foo: scan: %v", err)
	}
	want := result{ID: 1, K: "key", V: "value"}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("result of select (-want +got):\n%s", diff)
	}
}
