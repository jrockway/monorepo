package main

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"testing"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/google/go-cmp/cmp"
	joonix "github.com/joonix/log"
	"github.com/jrockway/monorepo/jlog/pkg/parse"
	aurora "github.com/logrusorgru/aurora/v3"
	"github.com/sirupsen/logrus" // nolint:depguard
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ignoreTimeFormatter struct {
	*parse.DefaultOutputFormatter
	i int
}

func (f *ignoreTimeFormatter) FormatTime(s *parse.State, t time.Time, w parse.Buffer) {
	f.i++
	w.WriteString(strconv.Itoa(f.i)) //nolint:errcheck
	// Check that the time is in some plausible range.  These could be very close to time.Now(),
	// but for Bunyan we hard-coded some times that get farther into the past every day.
	if t.After(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)) && t.Before(time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)) {
		w.WriteString(" ok") //nolint:errcheck
	} else {
		w.WriteString(" fail") //nolint:errcheck
	}
}

func TestLoggers(t *testing.T) {
	exampleObject := map[string]interface{}{"foo": "bar"}
	exampleError := errors.New("whoa")
	testData := []struct {
		name              string
		skip              string
		wantMessagePrefix string
		wantLineSuffix    string
		ins               *parse.InputSchema
		f                 func(buf *bytes.Buffer)
	}{
		{
			name: "zap",
			ins: &parse.InputSchema{
				LevelKey:    "level",
				MessageKey:  "msg",
				TimeKey:     "ts",
				LevelFormat: parse.DefaultLevelParser,
				TimeFormat:  parse.StrictUnixTimeParser,
				Strict:      true,
			},
			f: func(buf *bytes.Buffer) {
				enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
				core := zapcore.NewCore(enc, zapcore.Lock(zapcore.AddSync(buf)), zap.DebugLevel)
				l := zap.New(core)
				l.Info("line 1")
				l.Info("line 2", zap.String("string", "value"), zap.Int("int", 42), zap.Any("object", exampleObject))
				l.Error("line 3", zap.Error(exampleError))
			},
		},
		{
			name: "logrus/joonix",
			ins: &parse.InputSchema{
				LevelKey:    "severity",
				MessageKey:  "message",
				TimeKey:     "timestamp",
				LevelFormat: parse.DefaultLevelParser,
				TimeFormat:  parse.DefaultTimeParser,
				Strict:      true,
			},
			f: func(buf *bytes.Buffer) {
				l := &logrus.Logger{
					Out:       buf,
					Formatter: joonix.NewFormatter(),
					Level:     logrus.DebugLevel,
				}
				l.Info("line 1")
				l.WithField("string", "value").WithField("int", 42).WithField("object", exampleObject).Info("line 2")
				l.WithError(exampleError).Error("line 3")
			},
		},
		{
			name: "logrus/json",
			ins: &parse.InputSchema{
				LevelKey:    "level",
				MessageKey:  "msg",
				TimeKey:     "time",
				LevelFormat: parse.DefaultLevelParser,
				TimeFormat:  parse.DefaultTimeParser,
				Strict:      true,
			},
			f: func(buf *bytes.Buffer) {
				l := &logrus.Logger{
					Out:       buf,
					Formatter: new(logrus.JSONFormatter),
					Level:     logrus.DebugLevel,
				}
				l.Info("line 1")
				l.WithField("string", "value").WithField("int", 42).WithField("object", exampleObject).Info("line 2")
				l.WithError(exampleError).Error("line 3")
			},
		},
		{
			name: "logrus/json-stackdriver",
			ins: &parse.InputSchema{
				LevelKey:    "severity",
				MessageKey:  "message",
				TimeKey:     "time",
				LevelFormat: parse.DefaultLevelParser,
				TimeFormat:  parse.DefaultTimeParser,
				Strict:      true,
			},
			f: func(buf *bytes.Buffer) {
				l := &logrus.Logger{
					Out: buf,
					Formatter: &logrus.JSONFormatter{
						FieldMap: logrus.FieldMap{
							logrus.FieldKeyTime:  "time",
							logrus.FieldKeyLevel: "severity",
							logrus.FieldKeyMsg:   "message",
						},
						TimestampFormat: time.RFC3339Nano,
					},
					Level: logrus.DebugLevel,
				}
				l.Info("line 1")
				l.WithField("string", "value").WithField("int", 42).WithField("object", exampleObject).Info("line 2")
				l.WithError(exampleError).Error("line 3")
			},
		},
		{
			name: "lager/pretty",
			ins: &parse.InputSchema{
				LevelKey:    "level",
				MessageKey:  "message",
				TimeKey:     "timestamp",
				LevelFormat: parse.DefaultLevelParser,
				TimeFormat:  parse.DefaultTimeParser,
				Strict:      true,
				UpgradeKeys: []string{"data"},
			},
			wantMessagePrefix: "test.",
			wantLineSuffix:    " source:test",
			f: func(buf *bytes.Buffer) {
				l := lager.NewLogger("test")
				l.RegisterSink(lager.NewPrettySink(buf, lager.DEBUG))
				l.Info("line 1")
				l.Info("line 2", lager.Data{"string": "value", "int": 42, "object": exampleObject})
				l.Error("line 3", exampleError)
			},
		},
		{
			name: "lager",
			ins: &parse.InputSchema{
				LevelKey:    "log_level",
				MessageKey:  "message",
				TimeKey:     "timestamp",
				LevelFormat: parse.LagerLevelParser,
				TimeFormat:  parse.StrictUnixTimeParser,
				Strict:      true,
				UpgradeKeys: []string{"data"},
			},
			wantMessagePrefix: "test.",
			wantLineSuffix:    " source:test",
			f: func(buf *bytes.Buffer) {
				l := lager.NewLogger("test")
				l.RegisterSink(lager.NewWriterSink(buf, lager.DEBUG))
				l.Info("line 1")
				l.Info("line 2", lager.Data{"string": "value", "int": 42, "object": exampleObject})
				l.Error("line 3", exampleError)
			},
		},
		{
			name: "bunyan",
			ins: &parse.InputSchema{
				LevelKey:    "level",
				MessageKey:  "msg",
				TimeKey:     "time",
				LevelFormat: parse.BunyanV0LevelParser,
				TimeFormat:  parse.DefaultTimeParser,
				Strict:      true,
				DeleteKeys:  []string{"v"},
			},
			f: func(buf *bytes.Buffer) {
				// This is a node library.  We could bundle webpack and a JS
				// interpreter, but I just ran this program and copied in the
				// output.

				// var bunyan = require("bunyan");
				// var log = bunyan.createLogger({ name: "test" });
				// log.info("line 1");
				// log.info({ string: "value", int: 42, object: { foo: "bar" } }, "line 2");
				// log.error({ error: "whoa" }, "line 3");
				buf.WriteString(`{"level":30,"msg":"line 1","time":"2021-03-09T17:44:26.203Z","v":0}
{"level":30,"string":"value","int":42,"object":{"foo":"bar"},"msg":"line 2","time":"2021-03-09T17:44:26.204Z","v":0}
{"level":50,"error":"whoa","msg":"line 3","time":"2021-03-09T17:44:26.204Z","v":0}
`)
			},
		},
	}

	f := &ignoreTimeFormatter{
		DefaultOutputFormatter: &parse.DefaultOutputFormatter{
			Aurora:               aurora.NewAurora(false),
			AbsoluteTimeFormat:   "",
			ElideDuplicateFields: false,
		},
	}
	outs := &parse.OutputSchema{
		PriorityFields: []string{"error", "string", "int", "object", "source"},
		Formatter:      f,
	}
	fs := &parse.FilterScheme{}
	golden := `
INFO  1 ok line 1
INFO  2 ok line 2 string:value int:42 object:{"foo":"bar"}
ERROR 3 ok line 3 error:whoa
`[1:]
	for _, test := range testData {
		subTests := map[string]*parse.InputSchema{
			test.name:            test.ins,
			test.name + "_guess": {Strict: true},
		}
		for name, ins := range subTests {
			t.Run(name, func(t *testing.T) {
				f.i = 0
				outs.EmitErrorFn = func(msg string) { t.Fatalf("EmitErrorFn: %s", msg) }
				input := new(bytes.Buffer)
				output := new(bytes.Buffer)
				test.f(input)
				inputCopy := *input
				if _, err := parse.ReadLog(input, output, ins, outs, fs); err != nil {
					t.Fatalf("readlog: %v", err)
				}
				want := golden
				if p := test.wantMessagePrefix; p != "" {
					for _, msg := range []string{"line 1", "line 2", "line 3"} {
						want = strings.Replace(want, msg, p+msg, 1)
					}
				}
				if s := test.wantLineSuffix; s != "" {
					want = strings.ReplaceAll(want, "\n", s+"\n")
				}
				if test.skip != "" {
					t.Logf("skipped test:\noutput:\n%s", output.String())
					t.Logf("skipped test:\ninput:\n---\n%s\n---\n", inputCopy.String())
					t.Skip(test.skip)
				}
				if diff := cmp.Diff(output.String(), want); diff != "" {
					t.Errorf("output:\n%s", diff)
					t.Logf("input:\n---\n%s\n---\n", inputCopy.String())
				}
			})
		}
	}
}
