package types

import "github.com/rs/zerolog"

// Logger encapsulates a logger with an associated log level and toggle
type Logger struct {
	Enabled bool
	Level   string
	Log     *zerolog.Logger
}

// Application defines the application api
type Application interface {
	Create(data any) any
	Delete(id string) any
	Detail(id string) any
	List(*RequestMeta) any
	Update(data any) any
}
