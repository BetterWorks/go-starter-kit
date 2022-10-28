package middleware

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/jasonsites/gosk-api/internal/core"
)

func RequestLogger(logger *core.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		trace := ctx.MustGet("trace").(Trace)

		log := logger.Log.With().Str("ip", ip).Str("requestId", trace.RequestID).Logger()
		ctx.Set("log", log)

		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Error().Err(err)
		}
		header, err := json.Marshal(ctx.Request.Header)
		if err != nil {
			log.Error().Err(err)
		}
		method := ctx.Request.Method
		url := ctx.Request.URL.String()

		if logger.Enabled {
			if logger.Level == "debug" || logger.Level == "trace" {
				log.Trace().RawJSON("body", body).RawJSON("headers", header).
					Str("method", method).Str("url", url).Send()
			} else {
				log.Info().RawJSON("headers", header).Str("method", method).Str("url", url).Send()
			}
		}

		ctx.Next()
	}
}
