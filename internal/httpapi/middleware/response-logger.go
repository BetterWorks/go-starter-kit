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

// ResponseAndLogWriter extends gin.ResponseWriter with a bytes.Buffer for the response body
type ResponseAndLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

// Write extends ResponseWriter.Write by first adding response body to the ResponseAndLogWriter.Body buffer
func (rw *ResponseAndLogWriter) Write(b []byte) (int, error) {
	rw.Body.Write(b)
	return rw.ResponseWriter.Write(b)
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
		rw := &ResponseAndLogWriter{
			Body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = rw

		ctx.Next()

		if logger.Enabled {
			reqLD := ctx.MustGet("RequestLogData").(RequestLogData)
			trace := ctx.MustGet("Trace").(Trace)
			log := logger.Log.With().Str("req_id", trace.RequestID).Logger()

			body := rw.Body.Bytes()
			headers, err := json.Marshal(rw.Header())
			if err != nil {
				log.Error().Err(err)
			}

			data := newResponseLogData(ctx, body, headers, reqLD.Start)
			event := newResponseLogEvent(data, logger.Level, log)
			event.Send()
		}
	}
}

// newResponseLogData
func newResponseLogData(ctx *gin.Context, body, headers []byte, start time.Time) ResponseLogData {
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
