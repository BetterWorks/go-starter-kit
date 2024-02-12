package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/BetterWorks/go-starter-kit/internal/core/app"
	cl "github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/core/trace"
	"github.com/go-chi/chi/v5/middleware"
)

// ResponseLogData defines the data captured for response logging
type ResponseLogData struct {
	Body         map[string]any
	BodySize     *int
	Headers      http.Header
	ResponseTime string
	Status       int
}

// ResponseLoggerConfig defines necessary components for the response logger middleware
type ResponseLoggerConfig struct {
	Logger *cl.CustomLogger `validate:"required"`
	Next   func(r *http.Request) bool
}

// ResponseLogger returns the response logger middleware
func ResponseLogger(c *ResponseLoggerConfig) func(http.Handler) http.Handler {
	if err := app.Validator.Validate.Struct(c); err != nil {
		panic(err)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if c.Next != nil && c.Next(r) {
				next.ServeHTTP(w, r)
				return
			}

			// mark response time start
			start := time.Now()

			// extended response writer
			extendedRW := middleware.NewWrapResponseWriter(w, 1)
			bodyBuffer := new(bytes.Buffer)
			extendedRW.Tee(bodyBuffer)

			// call next middleware
			next.ServeHTTP(extendedRW, r)

			// calc response time and set header
			elapsed := time.Since(start).Milliseconds()
			responseTime := fmt.Sprintf("%sms", strconv.FormatInt(int64(elapsed), 10))
			extendedRW.Header().Set("X-Response-Time", responseTime)

			lrc := logResponseConfig{
				bodyBuffer:   bodyBuffer,
				logger:       c.Logger,
				request:      r,
				response:     extendedRW,
				responseTime: responseTime,
			}

			if err := logResponse(lrc); err != nil {
				c.Logger.Log.Error(err.Error())
			}
		})
	}
}

type logResponseConfig struct {
	bodyBuffer   *bytes.Buffer
	logger       *cl.CustomLogger
	request      *http.Request
	response     middleware.WrapResponseWriter
	responseTime string
}

func logResponse(c logResponseConfig) error {
	if c.logger.Enabled {
		traceID := trace.GetTraceIDFromContext(c.request.Context())
		log := c.logger.CreateContextLogger(traceID)

		bodyBytes := c.bodyBuffer.Bytes()
		bodySize := c.response.BytesWritten()

		data := &ResponseLogData{
			Headers:      c.response.Header(),
			ResponseTime: c.responseTime,
			Status:       c.response.Status(),
		}

		var body map[string]any
		if bodySize > 0 {
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				return err
			}

			data.Body = body
			data.BodySize = &bodySize
		}

		attrs := responseLogAttrs(data, c.logger.Level)
		log.With(attrs...).Info("response")
	}

	return nil
}

func responseLogAttrs(data *ResponseLogData, level string) []any {
	k := cl.AttrKey

	attrs := []any{
		slog.Int(k.HTTP.Status, data.Status),
		slog.String(k.ResponseTime, data.ResponseTime),
	}

	if data.BodySize != nil {
		attrs = append(attrs, slog.Int(k.HTTP.BodySize, *data.BodySize))
	}

	if level == cl.LevelDebug {
		if data.Body != nil {
			attrs = append(attrs, k.HTTP.Body, data.Body)
		}
		if data.Headers != nil {
			attrs = append(attrs, k.HTTP.Headers, data.Headers)
		}
	}

	return attrs
}
