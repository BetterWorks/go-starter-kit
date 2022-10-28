package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/jasonsites/gosk-api/internal/core"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ResponseLogger(logger *core.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		log := ctx.MustGet("log").(zerolog.Logger)

		var buf bytes.Buffer
		tee := io.TeeReader(ctx.Request.Response.Body, &buf)
		body, err := io.ReadAll(tee)
		if err != nil {
			fmt.Printf("error reading response body %+v\n", err)
		}
		ctx.Request.Response.Body = io.NopCloser(&buf)

		header, err := json.Marshal(ctx.Request.Response.Header)
		if err != nil {
			log.Error().Err(err)
		}

		status := ctx.Request.Response.StatusCode

		if logger.Enabled {
			if logger.Level == "debug" || logger.Level == "trace" {
				log.Trace().Str("body", string(body)).RawJSON("headers", header).Int("status", status).Send()
			} else {
				log.Info().RawJSON("headers", header).Int("status", status).Send()
			}
		}
	}
}
