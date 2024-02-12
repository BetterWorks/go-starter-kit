package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/BetterWorks/go-starter-kit/internal/core/app"
	cl "github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/core/trace"
	"github.com/BetterWorks/go-starter-kit/internal/http/jsonio"
)

// RequestLogData defines the data captured for request logging
type RequestLogData struct {
	Body     map[string]any
	ClientIP string
	Headers  http.Header
	Method   string
	Path     string
}

// RequestLoggerConfig defines necessary components for the request logger middleware
type RequestLoggerConfig struct {
	Logger *cl.CustomLogger `validate:"required"`
	Next   func(r *http.Request) bool
}

// RequestLogger returns the request logger middleware
func RequestLogger(c *RequestLoggerConfig) func(http.Handler) http.Handler {
	if err := app.Validator.Validate.Struct(c); err != nil {
		panic(err)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if c.Next != nil && c.Next(r) {
				next.ServeHTTP(w, r)
				return
			}

			if err := logRequest(w, r, c.Logger); err != nil {
				c.Logger.Log.Error(err.Error())
				jsonio.EncodeError(w, r, err)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func logRequest(w http.ResponseWriter, r *http.Request, logger *cl.CustomLogger) error {
	if logger.Enabled {
		traceID := trace.GetTraceIDFromContext(r.Context())
		log := logger.CreateContextLogger(traceID)

		var body map[string]any
		if logger.Level == cl.LevelDebug {
			r.Body = http.MaxBytesReader(w, r.Body, int64(1048576))

			copy, err := io.ReadAll(r.Body)
			if err != nil {
				return err
			}

			if len(copy) > 0 {
				if err := json.Unmarshal(copy, &body); err != nil {
					return err
				}
			}

			r.Body = io.NopCloser(bytes.NewBuffer(copy))
		}

		data := &RequestLogData{
			Body:     body,
			ClientIP: r.RemoteAddr,
			Headers:  r.Header,
			Method:   r.Method,
			Path:     r.URL.Path,
		}
		attrs := requestLogAttrs(data, logger.Level)
		log.With(attrs...).Info("request")
	}

	return nil
}

func requestLogAttrs(data *RequestLogData, level string) []any {
	k := cl.AttrKey

	attrs := []any{
		slog.String(k.IP, data.ClientIP),
		slog.String(k.HTTP.Method, data.Method),
		slog.String(k.HTTP.Path, data.Path),
	}

	if level == cl.LevelDebug {
		if data.Headers != nil {
			attrs = append(attrs, k.HTTP.Headers, data.Headers)
		}
		if data.Body != nil {
			attrs = append(attrs, k.HTTP.Body, data.Body)
		}
	}

	return attrs
}
