package middleware

import (
	"github.com/gin-gonic/gin"
)

func LocalType(t string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("type", t)
		ctx.Next()
	}
}
