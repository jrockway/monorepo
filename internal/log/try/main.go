package main

import (
	"time"

	"github.com/jrockway/monorepo/internal/log"
	"github.com/jrockway/monorepo/internal/pctx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Thing struct {
	Key   string `json:"k"`
	Value string `json:"v"`
}

func (t Thing) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("key", t.Key)
	enc.AddString("value", t.Value)
	return nil
}

func main() {
	log.InitLogger()
	log.SetLevel(log.DebugLevel)
	ctx := pctx.Background("")
	foo := zap.String("foo", "this is a long string")
	log.Debug(ctx, "this is a debug log", foo)
	log.Info(ctx, "this is an info log", foo, log.Proto("bytes", &wrapperspb.BytesValue{
		Value: []byte("these are some bytes that go on for a long time"),
	}))
	thing := zap.Object("thing", &Thing{Key: "key", Value: "value"})
	log.Error(ctx, "this is an error log", foo, zap.Duration("duration", 100*time.Millisecond), zap.Time("time", time.Now()), thing)

	// Make a child logger and race it with the parent, so that the race detector can examine
	// the locking.
	ch := make(chan struct{})
	child := log.ChildLogger(ctx, "child", log.WithFields(zap.Bool("in_child", true), zap.Namespace("child")))
	go func() {
		log.Info(child, "this is a child log message", foo, thing)
		ch <- struct{}{}
	}()
	go func() {
		time.Sleep(time.Millisecond)
		log.Info(ctx, "this is the original logger", foo, thing)
		ch <- struct{}{}
	}()
	for i := 0; i < 2; i++ {
		<-ch
	}
	child2 := log.ChildLogger(child, "child2", log.WithFields(zap.Bool("in_child2", true)))
	log.Info(child2, "hello from child2")
	log.Info(child, "hello from child1 again")
	log.Info(ctx, "hello from original logger again")
}
