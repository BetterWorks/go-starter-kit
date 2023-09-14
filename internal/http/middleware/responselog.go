package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/BetterWorks/gosk-api/internal/core/logger"
	"github.com/BetterWorks/gosk-api/internal/core/trace"
	"github.com/rs/zerolog"
)

// ExtendedResponseWriter extends gin.ResponseWriter with a bytes.Buffer to capture the response body
type ExtendedResponseWriter struct {
	http.ResponseWriter
	BodyLogBuffer *bytes.Buffer
}

// Write extends ResponseWriter.Write by first capturing response body to the ExtendedResponseWriter.BodyLogBuffer
func (erw *ExtendedResponseWriter) Write(b []byte) (int, error) {
	erw.BodyLogBuffer.Write(b)
	return erw.ResponseWriter.Write(b)
}

// ResponseLoggerConfig defines necessary components for the response logger middleware
type ResponseLoggerConfig struct {
	Logger *logger.Logger
	Next   func(r *http.Request) bool
}

// ResponseLogger returns the response logger middleware
func ResponseLogger(config *ResponseLoggerConfig) func(http.Handler) http.Handler {
	conf := setResponseLoggerConfig(config)
	logger := conf.Logger

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if conf.Next != nil && conf.Next(r) {
				next.ServeHTTP(w, r)
				return
			}

			start := time.Now()

			erw := &ExtendedResponseWriter{
				BodyLogBuffer:  bytes.NewBufferString(""),
				ResponseWriter: w,
			}
			w = erw

			// TODO: ensure upstream errors are not masked by possible log errors (json marshalling)
			next.ServeHTTP(w, r)

			rt := calcResponseTime(start)
			erw.Header().Set("X-Response-Time", rt)

			if err := logResponse(erw, r, logger, rt); err != nil {
				logger.Log.Error().Err(err)
				return
			}
		})
	}
}

// setResponseLoggerConfig returns a valid response logger configuration
func setResponseLoggerConfig(c *ResponseLoggerConfig) *ResponseLoggerConfig {
	if c.Logger == nil {
		log.Panicf("request logger middleware missing logger configuration")
	}

	return c
}

// logResponse logs the captured response data
func logResponse(erw *ExtendedResponseWriter, r *http.Request, logger *logger.Logger, rt string) error {
	if logger.Enabled {
		traceID := trace.GetTraceIDFromContext(r.Context())
		log := logger.CreateContextLogger(traceID)

		body := erw.BodyLogBuffer.Bytes()
		headers, err := json.Marshal(erw.Header())
		if err != nil {
			log.Error().Err(err)
		}

		data := newResponseLogData(erw, r, body, headers, rt)
		event := newResponseLogEvent(data, logger.Level, log)
		event.Send()
	}

	return nil
}

// ResponseLogData defines the data captured for response logging
type ResponseLogData struct {
	Body         []byte
	BodySize     int
	Headers      []byte
	ResponseTime string
	Status       int
}

func calcResponseTime(start time.Time) string {
	elapsed := time.Since(start).Milliseconds()
	return strconv.FormatInt(int64(elapsed), 10) + "ms"
}

// newResponseLogData captures relevant data from the response
func newResponseLogData(w *ExtendedResponseWriter, r *http.Request, body, headers []byte, rt string) *ResponseLogData {

	return &ResponseLogData{
		Body:         body,
		BodySize:     len(body),
		Headers:      headers,
		ResponseTime: rt,
		Status:       200, // TODO
	}
}

// newResponseLogEvent composes a new sendable response log event
func newResponseLogEvent(data *ResponseLogData, level string, log zerolog.Logger) *zerolog.Event {
	event := log.Info().
		Int("status", data.Status).
		Str("response_time", data.ResponseTime).
		RawJSON("headers", data.Headers).
		Int("body_size", data.BodySize)

	if level == "debug" || level == "trace" {
		if data.Body != nil {
			event.RawJSON("body", data.Body)
		}
	}

	return event
}
