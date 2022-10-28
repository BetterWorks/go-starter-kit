package core

import "github.com/rs/zerolog"

// Logger encapsulates a logger with an associated log level and toggle
type Logger struct {
	Enabled bool
	Level   string
	Log     *zerolog.Logger
}
