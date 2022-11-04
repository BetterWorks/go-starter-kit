package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jasonsites/gosk-api/internal/core/types"
	"github.com/rs/zerolog"
)

// RequestLogData
type RequestLogData struct {
	Body     string
	ClientIP string
	Headers  []byte
	Method   string
	Path     string
	Start    time.Time
}

// RequestLogger
func RequestLogger(logger *types.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if logger.Enabled {
			trace := ctx.MustGet("Trace").(types.Trace)
			log := logger.Log.With().Str("req_id", trace.RequestID).Logger()

			bodyBuf, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				log.Error().Err(err)
			}
			bodyRC1 := io.NopCloser(bytes.NewBuffer(bodyBuf))
			bodyRC2 := io.NopCloser(bytes.NewBuffer(bodyBuf))

			body := readBody(bodyRC1)
			ctx.Request.Body = bodyRC2

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

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.String()
}

// newRequestLogData
func newRequestLogData(ctx *gin.Context, body string, headers []byte) RequestLogData {
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
		event.Str("body", data.Body)
	}

	return event
}
