package log

import (
	"context"
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func BenchmarkFields(b *testing.B) {
	ctx, w := NewBenchLogger(false)
	for i := 0; i < b.N; i++ {
		Debug(ctx, "debug", zap.Int("i", i))
	}
	if w.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}

func BenchmarkFieldsSampled(b *testing.B) {
	ctx, w := NewBenchLogger(true)
	for i := 0; i < b.N; i++ {
		Debug(ctx, "debug", zap.Int("i", i))
	}
	if w.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}

func BenchmarkSpan(b *testing.B) {
	ctx, w := NewBenchLogger(false)
	errEven := errors.New("even")
	for i := 0; i < b.N; i++ {
		func() (retErr error) {
			defer Span(ctx, "bench", zap.Int("i", i))()
			if i%2 == 0 {
				return errEven //nolint:wrapcheck
			}
			return nil
		}() //nolint:errcheck
	}
	if w.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}

func BenchmarkSpanWithError(b *testing.B) {
	ctx, w := NewBenchLogger(false)
	errEven := errors.New("even")
	for i := 0; i < b.N; i++ {
		func() (retErr error) {
			defer Span(ctx, "bench", zap.Int("i", i))(Errorp(&retErr))
			if i%2 == 0 {
				return errEven //nolint:wrapcheck
			}
			return nil
		}() //nolint:errcheck
	}
	if w.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}

func BenchmarkLogrusFields(b *testing.B) {
	l := logrus.New()
	w := new(byteCounter)
	l.Formatter = &logrus.JSONFormatter{}
	l.Out = w
	l.Level = logrus.DebugLevel
	for i := 0; i < b.N; i++ {
		l.WithField("i", i).Debug("debug")
	}
	if w.c.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}

func BenchmarkLogrusSugar(b *testing.B) {
	l := logrus.New()
	w := new(byteCounter)
	l.Formatter = &logrus.JSONFormatter{}
	l.Out = w
	l.Level = logrus.DebugLevel
	for i := 0; i < b.N; i++ {
		l.Debugf("debug: %d", i)
	}
	if w.c.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}

func BenchmarkLogrusWrapper(b *testing.B) {
	ctx, w := NewBenchLogger(false)
	l := NewLogrus(ctx)
	for i := 0; i < b.N; i++ {
		l.Debugf("debug: %d", i)
	}
	if w.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}

func BenchmarkContextInfo_Logged(b *testing.B) {
	ctx, w := NewBenchLogger(false)
	dctx, c := context.WithCancelCause(ctx)
	for i := 0; i < b.N/2; i++ {
		Debug(dctx, "this is a log message")
	}
	c(errors.New("we are done here"))
	for i := b.N / 2; i < b.N; i++ {
		Debug(dctx, "this is a log message from an expired context")
	}
	if w.Load() == 0 {
		b.Fatal("no bytes added to logger")
	}
}

func BenchmarkContextInfo_NotLogged(b *testing.B) {
	ctx, w := newBenchInfoLogger(false)
	dctx, c := context.WithCancelCause(ctx)
	for i := 0; i < b.N/2; i++ {
		Debug(dctx, "this is a log message")
	}
	c(errors.New("we are done here"))
	for i := b.N / 2; i < b.N; i++ {
		Debug(dctx, "this is a log message from an expired context")
	}
	if w.Load() != 0 {
		b.Fatal("bytes unexpectedly added to logger")
	}
}
