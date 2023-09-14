package logger

import (
	"github.com/BetterWorks/gosk-api/internal/core/trace"
	"github.com/rs/zerolog"
)

// Logger encapsulates a logger with an associated log level and toggle
type Logger struct {
	Enabled bool
	Level   string
	Log     *zerolog.Logger
}

// CreateContextLogger returns a new child logger with attached trace ID
func (l *Logger) CreateContextLogger(traceID string) zerolog.Logger {
	return l.Log.With().Str(string(trace.TraceIDContextKey), traceID).Logger()
}
