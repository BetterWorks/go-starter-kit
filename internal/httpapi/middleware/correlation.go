package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/types"
)

// TODO https://opentracing.io/

// CorrelationContextKey
var CorrelationContextKey string

// Correlation
func Correlation(config *CorrelationConfig) fiber.Handler {
	conf := setCorrelationConfig(config)

	return func(ctx *fiber.Ctx) error {
		if conf.Next != nil && conf.Next(ctx) {
			return ctx.Next()
		}

		headers := ctx.GetReqHeaders()

		requestID := ctx.Get(conf.Header, conf.Generator())
		if headers[conf.Header] == "" {
			headers[conf.Header] = requestID
		}
		ctx.Set(conf.Header, requestID)

		trace := types.Trace{
			Headers:   headers,
			RequestID: requestID,
		}
		ctx.Locals(conf.ContextKey, &trace)

		return ctx.Next()
	}
}

// CorrelationConfig
type CorrelationConfig struct {
	// ContextKey for storing Trace data in context locals
	ContextKey string

	// Generator defines a function to generate request identifier
	Generator func() string

	// Header key for request ID get/set
	Header string

	// Next defines a function to skip this middleware on return true
	Next func(c *fiber.Ctx) bool
}

// setCorrelationConfig sets default CorrelationConfig values and CorrelationContextKey
func setCorrelationConfig(c *CorrelationConfig) *CorrelationConfig {
	var conf = &CorrelationConfig{
		ContextKey: "Trace",
		Generator:  uuid.NewString,
		Header:     "X-Request-ID",
		Next:       nil,
	}
	// override defaults
	if c.Header != "" {
		conf.Header = c.Header
	}
	if c.Generator != nil {
		conf.Generator = c.Generator
	}
	if c.ContextKey != "" {
		conf.ContextKey = c.ContextKey
	}

	// set exposed context key for use in other handlers
	CorrelationContextKey = conf.ContextKey

	return conf
}
