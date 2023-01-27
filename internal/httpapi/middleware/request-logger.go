package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jasonsites/gosk-api/internal/types"
	"github.com/rs/zerolog"
)

// RequestLoggerContextKey
var RequestLoggerContextKey string

// RequestLogger
func RequestLogger(config *RequestLoggerConfig) fiber.Handler {
	conf := setRequestLoggerConfig(config)
	logger := conf.Logger

	return func(ctx *fiber.Ctx) error {
		if conf.Next != nil && conf.Next(ctx) {
			return ctx.Next()
		}

		if logger.Enabled {
			requestID := ctx.Locals(types.CorrelationContextKey).(*types.Trace).RequestID
			log := logger.Log.With().Str("req_id", requestID).Logger()

			var body []byte
			if len(ctx.Body()) > 0 {
				b := new(bytes.Buffer)
				if err := json.Compact(b, ctx.Body()); err != nil {
					// log.Error().Err(err)
					return err
				}
				body = b.Bytes()
			}

			headers, err := json.Marshal(ctx.GetReqHeaders())
			if err != nil {
				// log.Error().Err(err)
				return err
			}

			data := newRequestLogData(ctx, body, headers)
			ctx.Locals("RequestLogData", data)

			event := newRequestLogEvent(data, logger.Level, log)
			event.Send()
		}

		return ctx.Next()
	}
}

// RequestLoggerConfig
type RequestLoggerConfig struct {
	ContextKey string
	Logger     *types.Logger
	Next       func(c *fiber.Ctx) bool
}

// setRequestLoggerConfig
func setRequestLoggerConfig(c *RequestLoggerConfig) *RequestLoggerConfig {
	if c.Logger == nil {
		log.Panicf("request logger middleware missing logger configuration")
	}
	conf := c

	// override defaults
	if c.ContextKey == "" {
		conf.ContextKey = "RequestLogData"
	}
	// set exposed context key for use in other handlers
	RequestLoggerContextKey = conf.ContextKey

	return conf
}

// RequestLogData
type RequestLogData struct {
	Body     []byte
	ClientIP string
	Headers  []byte
	Method   string
	Path     string
	Start    time.Time
}

// newRequestLogData
func newRequestLogData(ctx *fiber.Ctx, body, headers []byte) *RequestLogData {
	return &RequestLogData{
		Body:     body,
		ClientIP: ctx.IP(),
		Headers:  headers,
		Method:   ctx.Method(),
		Path:     ctx.Path(),
		Start:    time.Now(),
	}
}

// newRequestLogEvent
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
