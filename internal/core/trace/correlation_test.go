package trace

import (
	"context"
	"testing"

	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	fake "github.com/brianvoe/gofakeit/v6"
)

func Test_CreateOpContext(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ctx  context.Context
	}{{
		name: "with context provided",
		ctx:  context.Background(),
	}, {
		name: "without context provided",
		ctx:  nil,
	}}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			traceID := fake.UUID()
			context := CreateOpContext(tc.ctx, traceID)
			result := context.Value(TraceIDContextKey)

			if result != traceID {
				te.NewLineErrorf(t, traceID, result)
			}
		})
	}
}

func Test_GetTraceIDFromContext(t *testing.T) {
	t.Parallel()

	traceID := fake.UUID()

	tests := []struct {
		name            string
		traceID         string
		expectedTraceID string
	}{{
		name:            "with a traceID",
		traceID:         traceID,
		expectedTraceID: string(traceID),
	}, {
		name:            "without a traceID",
		traceID:         "",
		expectedTraceID: "unknown",
	}}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			context := CreateOpContext(context.Background(), tc.traceID)
			result := GetTraceIDFromContext(context)

			if result != tc.traceID {
				te.NewLineErrorf(t, tc.traceID, result)
			}
		})
	}
}
