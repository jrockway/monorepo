// Command postgres starts a postgres server.
package main

import (
	"fmt"

	"github.com/jrockway/monorepo/internal/log"
	"github.com/jrockway/monorepo/internal/pctx"
	"github.com/jrockway/monorepo/internal/testpostgres"
	"go.uber.org/zap"
)

func main() {
	log.InitLogger()
	log.SetLevel(log.DebugLevel)
	ctx, c := pctx.Interactive()
	defer c()
	cfg, cleanup, err := testpostgres.RunPostgres(ctx)
	if err != nil {
		log.Exit(ctx, "problem starting postgres", zap.Error(err))
	}
	log.Info(ctx, "running until Control-C", zap.String("connString", cfg.ConnString()))
	fmt.Println(cfg.ConnString())
	<-ctx.Done()
	log.Info(ctx, "awaiting cleanup")
	cleanup()
}
