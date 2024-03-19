package logger

import (
	"context"
	"log/slog"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/trace"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	fake "github.com/brianvoe/gofakeit/v6"
)

type CustomHandler struct {
	Attrs []slog.Attr
}

func (h CustomHandler) Enabled(c context.Context, l slog.Level) bool {
	return true
}

func (h CustomHandler) Handle(c context.Context, r slog.Record) error {
	return nil
}

func (h CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.Attrs = append(h.Attrs, attrs...)
	return h
}

func (h CustomHandler) WithGroup(name string) slog.Handler {
	return h
}

func Test_CreateContextLogger(t *testing.T) {
	t.Parallel()

	logger := CustomLogger{
		Enabled: false,
		Level:   LevelDebug,
		Log:     slog.New(CustomHandler{}),
	}

	traceID := fake.UUID()

	result := logger.CreateContextLogger(traceID)
	resultHandler := result.Handler().(CustomHandler)

	if result == nil {
		t.Errorf("Expected a logger but got nil")
	}

	if len(resultHandler.Attrs) != 1 {
		te.NewLineErrorf(t, 1, len(resultHandler.Attrs))
	}

	if resultHandler.Attrs[0].Key != string(trace.TraceIDContextKey) {
		te.NewLineErrorf(t, string(trace.TraceIDContextKey), resultHandler.Attrs[0].Key)
	}

	if resultHandler.Attrs[0].Value.String() != traceID {
		te.NewLineErrorf(t, traceID, resultHandler.Attrs[0].Value.String())
	}
}
