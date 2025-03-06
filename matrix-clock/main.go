package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/jrockway/monorepo/internal/log"
	"github.com/jrockway/monorepo/internal/pctx"
	"go.uber.org/zap"
	"golang.org/x/net/trace"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

func main() {
	log.InitLogger()
	log.SetLevel(log.DebugLevel)
	ctx := pctx.Background("")

	trace.AuthRequest = func(req *http.Request) (any bool, sensitive bool) {
		return true, true
	}

	if _, err := host.Init(); err != nil {
		log.Exit(ctx, "failed to init periph.io host", zap.Error(err))
	}

	i2cBus, err := i2creg.Open("2")
	if err != nil {
		log.Error(ctx, "problem opening i2c device", zap.Error(err))
		log.Info(ctx, "not monitoring sensors")
	}
	if i2cBus != nil {
		if err := monitorSensors(ctx, i2cBus); err != nil {
			log.Exit(ctx, "failed to init sensors", zap.Error(err))
		}
	}

	go drawClock(ctx)
	go watchGpsd(ctx)
	go watchChrony(ctx)

	http.HandleFunc("/", ServeStatus)

	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: http.DefaultServeMux,
	}
	log.AddLoggerToHTTPServer(ctx, "http", s)
	log.Info(ctx, "listening on :8080")
	if err := s.ListenAndServe(); err != nil {
		log.Exit(ctx, "http server failed", zap.Error(err))
	}
	log.Info(ctx, "exiting")
}
