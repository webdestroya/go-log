package tieredlog

import (
	"errors"
	"strings"
)

// ErrInvalidLevel is returned if the severity level is invalid.
var ErrInvalidLevel = errors.New("invalid level")

// Level of severity.
type Level int8

// Log levels.
const (
	InvalidLevel Level = iota - 1
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

var levelNames = [...]string{
	TraceLevel: "trace",
	DebugLevel: "debug",
	InfoLevel:  "info",
	WarnLevel:  "warn",
	ErrorLevel: "error",
	FatalLevel: "fatal",
}

var levelStrings = map[string]Level{
	"trace":   TraceLevel,
	"debug":   DebugLevel,
	"info":    InfoLevel,
	"warn":    WarnLevel,
	"warning": WarnLevel,
	"error":   ErrorLevel,
	"fatal":   FatalLevel,
}

// String implementation.
func (l Level) String() string {
	return levelNames[l]
}

// ParseLevel parses level string.
func ParseLevel(s string) (Level, error) {
	l, ok := levelStrings[strings.ToLower(s)]
	if !ok {
		return InvalidLevel, ErrInvalidLevel
	}

	return l, nil
}

// MustParseLevel parses level string or panics.
func MustParseLevel(s string) Level {
	l, err := ParseLevel(s)
	if err != nil {
		panic("invalid log level")
	}

	return l
}
