package log

import (
	"go.uber.org/zap/zapcore"
)

var (
	productionEncoder = zapcore.EncoderConfig{
		// These are important to keep consistent throughout versions.
		TimeKey:     "time",
		EncodeTime:  zapcore.RFC3339NanoTimeEncoder,
		LevelKey:    "severity",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		MessageKey:  "message",
		// These don't matter as much.
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// This is a less chatty console encoder.
	minimalConsoleEncoder = zapcore.EncoderConfig{
		TimeKey:          zapcore.OmitKey,
		LevelKey:         "L",
		NameKey:          "N",
		CallerKey:        "C",
		FunctionKey:      zapcore.OmitKey,
		MessageKey:       "M",
		StacktraceKey:    "S",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " ",
	}
)
