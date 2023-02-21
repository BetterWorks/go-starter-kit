package middleware

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jasonsites/gosk-api/internal/types"
	"github.com/rs/zerolog"
)

// ResponseLoggerConfig defines necessary components for the response logger middleware
type ResponseLoggerConfig struct {
	Logger *types.Logger
	Next   func(c *fiber.Ctx) bool
}

// ResponseLogger returns the response logger middleware
func ResponseLogger(config *ResponseLoggerConfig) fiber.Handler {
	conf := setResponseLoggerConfig(config)
	logger := conf.Logger

	return func(ctx *fiber.Ctx) error {
		if conf.Next != nil && conf.Next(ctx) {
			return ctx.Next()
		}

		// TODO: ensure upstream errors are not masked by possible log errors (json marshalling)
		if err := ctx.Next(); err != nil {
			if logError := logResponse(ctx, logger); logError != nil {
				return logError
			}
			return err
		}

		if logError := logResponse(ctx, logger); logError != nil {
			return logError
		}

		return nil
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
func logResponse(ctx *fiber.Ctx, logger *types.Logger) error {
	if logger.Enabled {
		requestID := ctx.Locals(types.CorrelationContextKey).(*types.Trace).RequestID
		log := logger.Log.With().Str("req_id", requestID).Logger()

		var body []byte
		if len(ctx.Response().Body()) > 0 {
			body = ctx.Response().Body()
		}

		headers, err := json.Marshal(ctx.GetRespHeaders())
		if err != nil {
			log.Error().Err(err).Msg("error marshalling response headers")
			return err
		}

		data := newResponseLogData(ctx, body, headers)
		event := newResponseLogEvent(data, logger.Level, log)
		event.Send()
	}

	return nil
}

// responseLogData defines the data captured for response logging
type ResponseLogData struct {
	Body         []byte
	BodySize     int
	Headers      []byte
	ResponseTime string
	Status       int
}

// newResponseLogData captures relevant data from the response
func newResponseLogData(ctx *fiber.Ctx, body, headers []byte) *ResponseLogData {
	start := ctx.Locals("RequestLogData").(*RequestLogData).Start
	elapsed := time.Since(start).Milliseconds()
	rt := strconv.FormatInt(int64(elapsed), 10) + "ms"

	return &ResponseLogData{
		Body:         body,
		BodySize:     len(body),
		Headers:      headers,
		ResponseTime: rt,
		Status:       ctx.Response().StatusCode(),
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
