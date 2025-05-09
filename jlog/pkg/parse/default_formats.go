package parse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	aurora "github.com/logrusorgru/aurora/v3"
)

type DefaultOutputFormatter struct {
	Aurora aurora.Aurora // Controls the use of color.

	// If true, print ↑ for fields that have an identical value as the previous line.
	ElideDuplicateFields bool

	// The time.Format string to show times in, like time.RFC3339.  If empty, show relative
	// times since the time the program started.  (A minus sign indicates the past; positive
	// values are in the future, good for when you are following a log file.)
	AbsoluteTimeFormat string

	// If non-empty, print only the fractional seconds for log lines that occurred on the same
	// second as the previous line.  For example, if SecondsOnlyFormat is set to ".000":
	//
	// INFO 2020-01-02T03:04:05.123Z first event
	// INFO                    .456  next event
	// INFO                    .789  another event
	// INFO 2020-01-02T03:04:05.000Z a brand new second
	//
	// Decimals are only aligned by careful selection of AbsoluteTimeFormat and
	// SecondsOnlyFormat strings.  The algorithm does nothing smart.
	SubSecondsOnlyFormat string

	Zone            *time.Location      // Zone is the time zone to display the output in.
	HighlightFields map[string]struct{} // HighlightFields visually distinguishes the named fields.
}

var (
	programStartTime = time.Now()
)

func (f *DefaultOutputFormatter) FormatTime(s *State, t time.Time, w Buffer) {
	var out string
	switch {
	case t.IsZero():
		out = "???"
		for utf8.RuneCountInString(out) < s.timePadding {
			out = " " + out
		}
	case f.AbsoluteTimeFormat == "":
		rel := t.Sub(programStartTime)
		abs := rel
		if rel < 0 {
			abs = -rel
		}
		var p time.Duration
		switch {
		case abs < time.Microsecond:
			p = time.Nanosecond
		case abs < time.Millisecond:
			p = time.Microsecond
		case abs < time.Second:
			p = time.Millisecond
		default:
			p = time.Second
		}
		out = rel.Truncate(p).String()
	case f.SubSecondsOnlyFormat != "":
		last := s.lastTime.Truncate(time.Second)
		if t.Sub(last) < time.Second && t.UnixNano() >= last.UnixNano() {
			out = t.In(f.Zone).Format(f.SubSecondsOnlyFormat)
		} else {
			out = t.In(f.Zone).Format(f.AbsoluteTimeFormat)
		}
	default:
		out = t.In(f.Zone).Format(f.AbsoluteTimeFormat)
	}
	for utf8.RuneCountInString(out) < s.timePadding {
		out += " "
	}
	if l := utf8.RuneCountInString(out); l > s.timePadding {
		s.timePadding = l
	}
	w.WriteString(f.Aurora.Green(out).String()) //nolint:errcheck
	s.lastTime = t
}

func cleanupNewlines(msg string) string {
	msg = strings.ReplaceAll(msg, "\n", "↩")
	msg = strings.ReplaceAll(msg, "\r", "←")
	return msg
}

func (f *DefaultOutputFormatter) FormatMessage(s *State, msg string, highlight bool, w Buffer) {
	msg = cleanupNewlines(msg)
	if highlight {
		msg = f.Aurora.Inverse(msg).String()
	}
	w.WriteString(msg) //nolint:errcheck
}

func (f *DefaultOutputFormatter) FormatLevel(s *State, level Level, w Buffer) {
	var l aurora.Value
	switch level {
	case LevelTrace:
		l = f.Aurora.Gray(15, "TRACE")
	case LevelDebug:
		l = f.Aurora.Blue("DEBUG")
	case LevelInfo:
		l = f.Aurora.Cyan("INFO ")
	case LevelWarn:
		l = f.Aurora.Yellow("WARN ")
	case LevelError:
		l = f.Aurora.Red("ERROR")
	case LevelPanic:
		l = f.Aurora.Magenta("PANIC")
	case LevelDPanic:
		l = f.Aurora.Magenta("DPANI")
	case LevelFatal:
		l = f.Aurora.BgMagenta("FATAL")
	default:
		l = f.Aurora.Gray(15, "UNK  ")
	}
	w.WriteString(l.String()) //nolint:errcheck
}

func (f *DefaultOutputFormatter) FormatField(s *State, k string, v interface{}, w Buffer) {
	var highlight bool
	if f.HighlightFields != nil {
		_, highlight = f.HighlightFields[k]
	}

	if highlight {
		w.WriteString(f.Aurora.Yellow(k).String()) //nolint:errcheck
	} else {
		w.WriteString(f.Aurora.Gray(16, k).String()) //nolint:errcheck
	}
	w.WriteString(f.Aurora.Gray(16, ":").String()) //nolint:errcheck

	var value []byte
	switch x := v.(type) {
	case string:
		x = cleanupNewlines(x)
		value = []byte(x)
	default:
		var err error
		value, err = json.Marshal(v)
		if err != nil {
			panic(fmt.Sprintf("marshal value: %v", err))
		}
	}

	if f.ElideDuplicateFields {
		old, ok := s.lastFields[k]
		if ok && bytes.Equal(old, value) {
			value = []byte("↑")
		} else {
			s.lastFields[k] = value
		}
	}

	w.Write(value) //nolint:errcheck
}
