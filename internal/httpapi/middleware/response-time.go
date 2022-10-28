package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ResponseTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		elapsed := time.Since(start).Milliseconds()
		rt := strconv.FormatInt(int64(elapsed), 10) + "ms"
		log := ctx.MustGet("log").(zerolog.Logger)
		log.Info().Str("responseTime", rt).Send()
	}
}
