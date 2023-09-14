package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/BetterWorks/gosk-api/internal/core/logger"
	"github.com/BetterWorks/gosk-api/internal/core/trace"
	"github.com/rs/zerolog"
)

// RequestLogContextKey
var RequestLogContextKey trace.ContextKey

// RequestLoggerConfig defines necessary components for the request logger middleware
type RequestLoggerConfig struct {
	ContextKey trace.ContextKey
	Logger     *logger.Logger
	Next       func(r *http.Request) bool
}

// RequestLogger returns the request logger middleware
func RequestLogger(config *RequestLoggerConfig) func(http.Handler) http.Handler {
	conf := setRequestLoggerConfig(config)
	logger := conf.Logger

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if conf.Next != nil && conf.Next(r) {
				next.ServeHTTP(w, r)
				return
			}

			data, err := logRequest(w, r, logger)
			if err != nil {
				logger.Log.Error().Err(err)
			}

			ctx := context.WithValue(r.Context(), RequestLogContextKey, data)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// setRequestLoggerConfig returns a valid request logger configuration and sets the context key for external use
func setRequestLoggerConfig(c *RequestLoggerConfig) *RequestLoggerConfig {
	if c.Logger == nil {
		log.Panicf("request logger middleware missing logger configuration")
	}
	conf := c

	// override defaults
	if c.ContextKey == "" {
		conf.ContextKey = "request_data"
	}
	// set middleware-scoped context key for use in other handlers
	RequestLogContextKey = conf.ContextKey

	return conf
}

// logRequest logs the captured request data
func logRequest(w http.ResponseWriter, r *http.Request, logger *logger.Logger) (*RequestLogData, error) {
	if logger.Enabled {
		traceID := trace.GetTraceIDFromContext(r.Context())
		log := logger.CreateContextLogger(traceID)

		maxBytes := 1_048_576
		r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

		var body []byte
		n, err := r.Body.Read(body)
		if err != nil {
			return nil, err
		}

		if n > 0 {
			b := new(bytes.Buffer)
			if err := json.Compact(b, body); err != nil {
				log.Error().Err(err).Send()
				return nil, err
			}
			body = b.Bytes()
		}

		headers, err := json.Marshal(r.Header)
		if err != nil {
			log.Error().Err(err).Send()
			return nil, err
		}

		data := newRequestLogData(r, body, headers)
		event := newRequestLogEvent(data, logger.Level, log)
		event.Send()

		return data, nil
	}

	return nil, nil
}

// RequestLogData defines the data captured for request logging
type RequestLogData struct {
	Body     []byte
	ClientIP string
	Headers  []byte
	Method   string
	Path     string
}

// newRequestLogData captures relevant data from the request
func newRequestLogData(r *http.Request, body, headers []byte) *RequestLogData {
	return &RequestLogData{
		Body:     body,
		ClientIP: r.RemoteAddr,
		Headers:  headers,
		Method:   r.Method,
		Path:     r.URL.Path,
	}
}

// newRequestLogEvent composes a new sendable request log event
func newRequestLogEvent(data *RequestLogData, level string, log zerolog.Logger) *zerolog.Event {
	event := log.Info().
		Str("ip", data.ClientIP).
		Str("method", data.Method).
		Str("path", data.Path).
		RawJSON("headers", data.Headers)

	if level == "debug" || level == "trace" {
		if data.Body != nil {
			event.RawJSON("body", data.Body)
		}
	}

	return event
}
