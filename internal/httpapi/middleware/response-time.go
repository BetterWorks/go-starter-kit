package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func ResponseTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		elapsed := time.Since(start).Milliseconds()
		log.Printf("%v", elapsed)
	}
}
