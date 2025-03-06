package parse

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/jrockway/monorepo/internal/errors"
	"github.com/logrusorgru/aurora/v3"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var (
	zapBufferPool = buffer.NewPool()
)

type jlogEncoder struct {
	// Mutex protects against concurrent calls.  Cloned encoders share the jlog output state,
	// which isn't thread safe unless this mutex is copied.
	*sync.Mutex

	zapcore.Encoder                        // Encoder encodes fields to JSON using zap's exact algorithm.
	cfg             *zapcore.EncoderConfig // cfg lets us parse the encoded JSON and remove what jlog already handles.
	line            *line                  // line doesn't need to be allocated each time we call "emit".
	schema          *OutputSchema          // schema controls how output logs are displayed.
	fields          []string               // fields is the ordering of fields added to Encoder.
}

var _ zapcore.Encoder = (*jlogEncoder)(nil)

// ZapEncoderConfig configures the zap encoder.
type ZapEncoderConfig struct {
	NoColor    bool           // If true, don't colorize the output.
	NoElide    bool           // If true, don't elide fields that were the same as the line above.
	TimeFormat string         // Format of the time field.  If empty, default to hour:minute:second.millisecond.
	Zone       *time.Location // Time zone to display times in.  If nil, use the local time.
}

// NewJlogEncoder returns a zap encoder that outputs lines as though jlog was run on the JSON output.
func NewJlogEncoder(zcfg zapcore.EncoderConfig, jcfg ZapEncoderConfig) zapcore.Encoder {
	timeFormat := "15:04:05.000" // No day needed, user is reading the line as it's printed.
	if f := jcfg.TimeFormat; f != "" {
		timeFormat = f
	}
	zone := time.Local
	if z := jcfg.Zone; z != nil {
		zone = z
	}
	zcfg.MessageKey = ""
	zcfg.TimeKey = ""
	zcfg.LevelKey = ""
	f := &DefaultOutputFormatter{
		Aurora:               aurora.NewAurora(!jcfg.NoColor),
		AbsoluteTimeFormat:   timeFormat,
		Zone:                 zone,
		ElideDuplicateFields: !jcfg.NoElide,
	}
	l := new(line) // no need to allocate for each log
	return &jlogEncoder{
		Mutex:   new(sync.Mutex),
		Encoder: zapcore.NewJSONEncoder(zcfg),
		cfg:     &zcfg,
		line:    l,
		schema: &OutputSchema{
			Formatter: f,
			state: State{
				lastFields: make(map[string][]byte),
			},
		},
	}
}

// Clone implements zapcore.Encoder.
func (enc *jlogEncoder) Clone() zapcore.Encoder {
	result := &jlogEncoder{
		Mutex:   enc.Mutex,
		cfg:     enc.cfg,
		line:    new(line),
		schema:  enc.schema,
		Encoder: enc.Encoder.Clone(),
	}
	result.fields = append(result.fields, enc.fields...)
	return result
}

// EncodeEntry implements zapcore.Encoder.
func (enc *jlogEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	encoder := enc.Encoder
	enc.Lock()
	defer enc.Unlock()

	l := enc.line
	l.reset()
	switch entry.Level {
	case zapcore.DebugLevel:
		l.lvl = LevelDebug
	case zapcore.InfoLevel:
		l.lvl = LevelInfo
	case zapcore.WarnLevel:
		l.lvl = LevelWarn
	case zapcore.ErrorLevel:
		l.lvl = LevelError
	case zapcore.DPanicLevel:
		l.lvl = LevelDPanic
	case zapcore.PanicLevel:
		l.lvl = LevelPanic
	case zapcore.FatalLevel:
		l.lvl = LevelFatal
	default:
		l.lvl = LevelUnknown
	}
	l.msg = entry.Message
	l.time = entry.Time
	l.fields = make(map[string]any)

	// Jlog keeps field order consistent between multiple lines, but zap gives each line a
	// defined field ordering; use their algorithm instead of ours.  It's the least surprising
	// thing, I think.  Name, Caller, Function always come first; Stack always comes last.
	enc.schema.state.seenFields = []string{
		enc.cfg.NameKey,
		enc.cfg.CallerKey,
		enc.cfg.FunctionKey,
	}
	enc.schema.state.seenFields = append(enc.schema.state.seenFields, enc.fields...)
	for _, f := range fields {
		enc.schema.state.seenFields = append(enc.schema.state.seenFields, f.Key)
	}
	enc.schema.state.seenFields = append(enc.schema.state.seenFields, enc.cfg.StacktraceKey)

	// To marshal the fields, we encode using Zap's JSON encoder, and then unmarshal.  This most
	// closely mirrors the behavior of jlog, but is somewhat inefficient.  This strategy was chosen because of features
	js, err := encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, errors.Wrap(err, "mashal log to json")
	}
	defer js.Free()
	var decoded map[string]any
	if err := json.Unmarshal(js.Bytes(), &decoded); err != nil {
		return nil, errors.Wrap(err, "unmarhsal encoded log")
	}
	l.fields = decoded

	buf := zapBufferPool.Get()
	enc.schema.Emit(l, buf)
	return buf, nil
}

// These methods implement zapcore.Encoder, remembering the field name before passing through to the
// underlying encoder.  This lets us keep the field order consistent with the JSON output.

func (enc *jlogEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	enc.fields = append(enc.fields, key)
	return enc.Encoder.AddArray(key, marshaler) //nolint:wrapcheck
}
func (enc *jlogEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	enc.fields = append(enc.fields, key)
	return enc.Encoder.AddObject(key, marshaler) //nolint:wrapcheck
}
func (enc *jlogEncoder) AddBinary(key string, value []byte) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddBinary(key, value)
}
func (enc *jlogEncoder) AddByteString(key string, value []byte) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddByteString(key, value)
}
func (enc *jlogEncoder) AddBool(key string, value bool) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddBool(key, value)
}
func (enc *jlogEncoder) AddComplex128(key string, value complex128) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddComplex128(key, value)
}
func (enc *jlogEncoder) AddComplex64(key string, value complex64) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddComplex64(key, value)
}
func (enc *jlogEncoder) AddDuration(key string, value time.Duration) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddDuration(key, value)
}
func (enc *jlogEncoder) AddFloat64(key string, value float64) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddFloat64(key, value)
}
func (enc *jlogEncoder) AddFloat32(key string, value float32) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddFloat32(key, value)
}
func (enc *jlogEncoder) AddInt(key string, value int) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddInt(key, value)
}
func (enc *jlogEncoder) AddInt64(key string, value int64) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddInt64(key, value)
}
func (enc *jlogEncoder) AddInt32(key string, value int32) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddInt32(key, value)
}
func (enc *jlogEncoder) AddInt16(key string, value int16) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddInt16(key, value)
}
func (enc *jlogEncoder) AddInt8(key string, value int8) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddInt8(key, value)
}
func (enc *jlogEncoder) AddString(key, value string) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddString(key, value)
}
func (enc *jlogEncoder) AddTime(key string, value time.Time) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddTime(key, value)
}
func (enc *jlogEncoder) AddUint(key string, value uint) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddUint(key, value)
}
func (enc *jlogEncoder) AddUint64(key string, value uint64) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddUint64(key, value)
}
func (enc *jlogEncoder) AddUint32(key string, value uint32) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddUint32(key, value)
}
func (enc *jlogEncoder) AddUint16(key string, value uint16) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddUint16(key, value)
}
func (enc *jlogEncoder) AddUint8(key string, value uint8) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddUint8(key, value)
}
func (enc *jlogEncoder) AddUintptr(key string, value uintptr) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.AddUintptr(key, value)
}
func (enc *jlogEncoder) AddReflected(key string, value interface{}) error {
	enc.fields = append(enc.fields, key)
	return enc.Encoder.AddReflected(key, value) //nolint:wrapcheck
}
func (enc *jlogEncoder) OpenNamespace(key string) {
	enc.fields = append(enc.fields, key)
	enc.Encoder.OpenNamespace(key)
}
