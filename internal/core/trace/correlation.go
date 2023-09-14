package trace

import "context"

// Correlation
type Correlation struct {
	Headers map[string][]string
	TraceID ContextKey // TODO: consider uuid.UUID
}

// TraceIDContextKey defines the context key used for tracking operation trace ID
const TraceIDContextKey ContextKey = "trace_id"

// CreateOpContext creates an operation context with correlation data
func CreateOpContext(ctx context.Context, traceID string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ContextKey(string(TraceIDContextKey)), traceID)
}

// GetTraceIDFromContext retrieves the trace ID from the operation context
func GetTraceIDFromContext(ctx context.Context) string {
	val := ctx.Value(TraceIDContextKey)
	traceID, ok := val.(string)
	if !ok {
		traceID = "unknown"
	}
	return traceID
}
