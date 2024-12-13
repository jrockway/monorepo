package main

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"syscall"

	"github.com/jrockway/monorepo/internal/log"
	"github.com/jrockway/monorepo/internal/testpostgres"
	"go.uber.org/zap"
)

func main() {
	log.InitLogger()
	log.SetLevel(log.DebugLevel)
	ctx := log.AddLogger(context.Background())
	cmd, err := testpostgres.PsqlCmd(ctx)
	if err != nil {
		log.Exit(ctx, "problem resolving psql command", zap.Error(err))
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:    true,
		Foreground: true,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if len(os.Args) > 1 {
		cmd.Args = append(cmd.Args, os.Args[1:]...)
	}
	if err := cmd.Run(); err != nil {
		log.Error(ctx, "problem running psql", zap.Error(err))
		exit := new(exec.ExitError)
		if errors.As(err, &exit) {
			os.Exit(exit.ExitCode())
		}
	}
	os.Exit(0)
}
