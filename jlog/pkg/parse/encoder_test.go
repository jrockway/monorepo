package parse

import (
	"bytes"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ignorePanic(f func()) {
	defer func() {
		recover() //nolint:errcheck
	}()
	f()
}

type clock struct{}

func (clock) Now() time.Time {
	return time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
}

func (clock) NewTicker(time.Duration) *time.Ticker {
	return nil
}

func TestZapEncoder(t *testing.T) {
	enc := NewJlogEncoder(
		zapcore.EncoderConfig{
			NameKey:      "name",
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		ZapEncoderConfig{
			NoColor: true,
			Zone:    time.UTC,
		},
	)
	var buf bytes.Buffer
	l := zap.New(zapcore.NewCore(enc, zapcore.AddSync(&buf), zapcore.DebugLevel)).WithOptions(
		zap.AddCaller(),
		zap.WithPanicHook(zapcore.WriteThenPanic),
		zap.WithFatalHook(zapcore.WriteThenPanic),
		zap.WithClock(clock{}),
	).Named("root")

	fields := []zap.Field{
		zap.Int8("0", 8), zap.Int16("1", 16), zap.Int32("2", 32), zap.Int64("3", 64),
		zap.Uint8("4", 8), zap.Uint16("5", 16), zap.Uint32("6", 32), zap.Uint64("7", 64),
		zap.String("8", "string"), zap.Binary("9", []byte("bytes")),
		zap.Any("10", struct {
			Foo int
			Bar string
		}{42, "bar"}),
		zap.Any("11", []struct{ Foo int }{{1}, {2}, {3}}),
	}
	l.Debug("debug", fields...)
	l.Info("info", fields...)
	l.Warn("warn", fields...)
	l.Error("error", fields...)
	l.DPanic("dpanic", fields...)
	ignorePanic(func() { l.Panic("panic", fields...) })
	ignorePanic(func() { l.Fatal("fatal", fields...) })
	child := l.Named("child").With(zap.Bool("isChild", true))
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		child.Info("child info", zap.String("foo", "bar"), fields[0])
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Millisecond)
		l.Info("parent", zap.Bool("isChild", false), fields[0])
		wg.Done()
	}()
	wg.Wait()

	want := `DEBUG 00:00:00.000 debug name:root caller:parse/encoder_test.go:61 0:8 1:16 2:32 3:64 4:8 5:16 6:32 7:64 8:string 9:Ynl0ZXM= 10:{"Bar":"bar","Foo":42} 11:[{"Foo":1},{"Foo":2},{"Foo":3}]
INFO  00:00:00.000 info name:↑ caller:parse/encoder_test.go:62 0:↑ 1:↑ 2:↑ 3:↑ 4:↑ 5:↑ 6:↑ 7:↑ 8:↑ 9:↑ 10:↑ 11:↑
WARN  00:00:00.000 warn name:↑ caller:parse/encoder_test.go:63 0:↑ 1:↑ 2:↑ 3:↑ 4:↑ 5:↑ 6:↑ 7:↑ 8:↑ 9:↑ 10:↑ 11:↑
ERROR 00:00:00.000 error name:↑ caller:parse/encoder_test.go:64 0:↑ 1:↑ 2:↑ 3:↑ 4:↑ 5:↑ 6:↑ 7:↑ 8:↑ 9:↑ 10:↑ 11:↑
DPANI 00:00:00.000 dpanic name:↑ caller:parse/encoder_test.go:65 0:↑ 1:↑ 2:↑ 3:↑ 4:↑ 5:↑ 6:↑ 7:↑ 8:↑ 9:↑ 10:↑ 11:↑
PANIC 00:00:00.000 panic name:↑ caller:parse/encoder_test.go:66 0:↑ 1:↑ 2:↑ 3:↑ 4:↑ 5:↑ 6:↑ 7:↑ 8:↑ 9:↑ 10:↑ 11:↑
FATAL 00:00:00.000 fatal name:↑ caller:parse/encoder_test.go:67 0:↑ 1:↑ 2:↑ 3:↑ 4:↑ 5:↑ 6:↑ 7:↑ 8:↑ 9:↑ 10:↑ 11:↑
INFO  00:00:00.000 child info name:root.child caller:parse/encoder_test.go:72 isChild:true foo:bar 0:↑
INFO  00:00:00.000 parent name:root caller:parse/encoder_test.go:77 isChild:false 0:↑
`
	if diff := cmp.Diff(want, buf.String()); diff != "" {
		t.Errorf("output (-want +got):\n%s", diff)
	}
}
