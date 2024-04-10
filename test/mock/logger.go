package mock

import (
	"io"
	"log/slog"
)

// Logger returns a slog.Logger instance that writes to nothing
func Logger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
