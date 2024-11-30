// Package wrapcheck provides a wrapcheck analyzer.
package wrapcheck

import (
	"github.com/tomarrell/wrapcheck/v2/wrapcheck"
	"golang.org/x/tools/go/analysis"
)

var Analyzer *analysis.Analyzer

func init() {
	cfg := wrapcheck.NewDefaultConfig()
	cfg.IgnoreSigs = []string{
		"errors.New",
		"fmt.Errorf",
		"github.com/jrockway/monorepo/internal/errors.Errorf",
		"github.com/jrockway/monorepo/internal/errors.New",
		"github.com/jrockway/monorepo/internal/errors.Unwrap",
		"github.com/jrockway/monorepo/internal/errors.Close",
		"github.com/jrockway/monorepo/internal/errors.Invoke",
		"github.com/jrockway/monorepo/internal/errors.Invoke1",
		"google.golang.org/grpc/status.Error",
		"google.golang.org/grpc/status.Errorf",
		"(*google.golang.org/grpc/internal/status.Status).Err",
		"google.golang.org/protobuf/types/known/anypb.New",
		".Wrap(",
		".Wrapf(",
		".WithMessage(",
		".WithMessagef(",
		".WithStack(",
	}
	cfg.IgnorePackageGlobs = []string{
		"github.com/jrockway/monorepo/**",
	}
	cfg.IgnoreInterfaceRegexps = []string{}
	Analyzer = wrapcheck.NewAnalyzer(cfg)
}
