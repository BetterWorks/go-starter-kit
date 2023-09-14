package middleware

import (
	"net/http"

	"github.com/BetterWorks/gosk-api/internal/core/trace"
	"github.com/google/uuid"
)

// Correlation
func Correlation(config *CorrelationConfig) func(http.Handler) http.Handler {
	conf := setCorrelationConfig(config)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if conf.Next != nil && conf.Next(r) {
				next.ServeHTTP(w, r)
				return
			}

			headers := r.Header
			traceID := headers.Get(conf.Header)
			if traceID == "" {
				traceID = conf.Generator()
			}
			if headers[conf.Header] == nil {
				headers[conf.Header] = []string{traceID}
			}
			w.Header().Set(conf.Header, traceID)

			ctx := trace.CreateOpContext(r.Context(), traceID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// CorrelationConfig
type CorrelationConfig struct {
	// ContextKey for storing correlation data in context locals
	ContextKey trace.ContextKey

	// Generator defines a function to generate request identifier
	Generator func() string

	// Header key for trace ID get/set
	Header string

	// Next defines a function to skip this middleware on return true
	Next func(r *http.Request) bool
}

// setCorrelationConfig sets default CorrelationConfig values and CorrelationContextKey
func setCorrelationConfig(c *CorrelationConfig) *CorrelationConfig {
	// default config
	var conf = &CorrelationConfig{
		ContextKey: trace.TraceIDContextKey,
		Generator:  uuid.NewString,
		Header:     "X-Request-Id",
		Next:       nil,
	}

	// default overrides
	if c.ContextKey != "" {
		conf.ContextKey = c.ContextKey
	}
	if c.Generator != nil {
		conf.Generator = c.Generator
	}
	if c.Header != "" {
		conf.Header = c.Header
	}

	return conf
}
