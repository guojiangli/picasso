package klog

import "github.com/pkg/errors"

// Level defines log levels.
type Level int8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = -1
)

//LevelUnmarshalText level in config unmarshal to log.level
func LevelUnmarshalText(text string) (level Level, err error) {
	switch string(text) {
	case "debug", "DEBUG":
		return DebugLevel, nil
	case "info", "INFO", "": // make the zero value useful
		return InfoLevel, nil
	case "warn", "WARN":
		return WarnLevel, nil
	case "error", "ERROR":
		return ErrorLevel, nil
	case "panic", "PANIC":
		return PanicLevel, nil
	case "fatal", "FATAL":
		return FatalLevel, nil
	default:
		return DebugLevel, errors.New("unknown level")
	}
}
