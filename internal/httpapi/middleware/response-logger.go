package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jasonsites/gosk-api/internal/core/types"
	"github.com/rs/zerolog"
)

// ResponseLogger
func ResponseLogger(config *ResponseLoggerConfig) fiber.Handler {
	conf := setResponseLoggerConfig(config)
	logger := conf.Logger

	return func(ctx *fiber.Ctx) error {
		if conf.Next != nil && conf.Next(ctx) {
			return ctx.Next()
		}

		ctx.Next()

		if logger.Enabled {
			requestID := ctx.Locals(CorrelationContextKey).(*types.Trace).RequestID
			log := logger.Log.With().Str("req_id", requestID).Logger()

			var body []byte
			if len(ctx.Response().Body()) > 0 {
				body = ctx.Response().Body()
			}

			headers, err := json.Marshal(ctx.GetRespHeaders())
			if err != nil {
				// log.Error().Err(err)
				fmt.Printf("Error marshalling response headers: %+v\n", err)
				return err
			}

			data := newResponseLogData(ctx, body, headers)
			event := newResponseLogEvent(data, logger.Level, log)
			event.Send()
		}

		return nil
	}
}

// ResponseLoggerConfig
type ResponseLoggerConfig struct {
	Logger *types.Logger
	Next   func(c *fiber.Ctx) bool
}

// setResponseLoggerConfig
func setResponseLoggerConfig(c *ResponseLoggerConfig) *ResponseLoggerConfig {
	if c.Logger == nil {
		log.Panicf("request logger middleware missing logger configuration")
	}
	return c
}

// ResponseLogData defines the data captured for response logging
type ResponseLogData struct {
	Body         []byte
	BodySize     int
	Headers      []byte
	ResponseTime string
	Status       int
}

// newResponseLogData
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

// newResponseLogEvent
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
