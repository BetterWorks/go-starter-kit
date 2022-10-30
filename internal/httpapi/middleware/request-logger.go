package middleware

import (
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jasonsites/gosk-api/internal/core"
	"github.com/rs/zerolog"
)

// RequestLogData
type RequestLogData struct {
	Body     []byte
	ClientIP string
	Headers  []byte
	Method   string
	Path     string
	Start    time.Time
}

// RequestLogger
func RequestLogger(logger *core.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if logger.Enabled {
			trace := ctx.MustGet("Trace").(Trace)
			log := logger.Log.With().Str("req_id", trace.RequestID).Logger()

			body, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				log.Error().Err(err)
			}
			headers, err := json.Marshal(ctx.Request.Header)
			if err != nil {
				log.Error().Err(err)
			}

			data := newRequestLogData(ctx, body, headers)
			ctx.Set("RequestLogData", data)

			event := newRequestLogEvent(data, logger.Level, log)
			event.Send()
		}

		ctx.Next()
	}
}

// newRequestLogData
func newRequestLogData(ctx *gin.Context, body, headers []byte) RequestLogData {
	return RequestLogData{
		Body:     body,
		ClientIP: ctx.ClientIP(),
		Headers:  headers,
		Method:   ctx.Request.Method,
		Path:     ctx.Request.URL.String(),
		Start:    time.Now(),
	}
}

// newRequestLogEvent
func newRequestLogEvent(data RequestLogData, level string, log zerolog.Logger) *zerolog.Event {
	event := log.Info().
		Str("ip", data.ClientIP).
		Str("method", data.Method).
		Str("path", data.Path).
		RawJSON("headers", data.Headers)

	if level == "debug" || level == "trace" {
		event.RawJSON("body", data.Body)
	}

	return event
}
