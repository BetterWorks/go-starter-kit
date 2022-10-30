package middleware

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jasonsites/gosk-api/internal/core"
	"github.com/rs/zerolog"
)

// ExtendedResponseWriter extends gin.ResponseWriter with a bytes.Buffer to capture the response body
type ExtendedResponseWriter struct {
	gin.ResponseWriter
	BodyLogBuffer *bytes.Buffer
}

// Write extends ResponseWriter.Write by first capturing response body to the ExtendedResponseWriter.BodyLogBuffer
func (erw *ExtendedResponseWriter) Write(b []byte) (int, error) {
	erw.BodyLogBuffer.Write(b)
	return erw.ResponseWriter.Write(b)
}

// ResponseLogData defines the data captured for response logging
type ResponseLogData struct {
	Body         []byte
	BodySize     int
	Headers      []byte
	ResponseTime string
	Status       int
}

// ResponseLogger
func ResponseLogger(logger *core.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		erw := &ExtendedResponseWriter{
			BodyLogBuffer:  bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = erw

		ctx.Next()

		if logger.Enabled {
			trace := ctx.MustGet("Trace").(Trace)
			log := logger.Log.With().Str("req_id", trace.RequestID).Logger()

			body := erw.BodyLogBuffer.Bytes()
			headers, err := json.Marshal(erw.Header())
			if err != nil {
				log.Error().Err(err)
			}

			data := newResponseLogData(ctx, body, headers)
			event := newResponseLogEvent(data, logger.Level, log)
			event.Send()
		}
	}
}

// newResponseLogData
func newResponseLogData(ctx *gin.Context, body, headers []byte) ResponseLogData {
	start := ctx.MustGet("RequestLogData").(RequestLogData).Start
	elapsed := time.Since(start).Milliseconds()
	rt := strconv.FormatInt(int64(elapsed), 10) + "ms"

	return ResponseLogData{
		Body:         body,
		BodySize:     len(body),
		Headers:      headers,
		ResponseTime: rt,
		Status:       ctx.Writer.Status(),
	}
}

// newResponseLogEvent
func newResponseLogEvent(data ResponseLogData, level string, log zerolog.Logger) *zerolog.Event {
	event := log.Info().
		Int("status", data.Status).
		Str("response_time", data.ResponseTime).
		RawJSON("headers", data.Headers).
		Int("body_size", data.BodySize)

	if level == "debug" || level == "trace" {
		event.RawJSON("body", data.Body)
	}

	return event
}
