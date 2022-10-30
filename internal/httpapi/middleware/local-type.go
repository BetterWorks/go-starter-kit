package middleware

import (
	"github.com/gin-gonic/gin"
)

// LocalType
func LocalType(t string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("Type", t)
		ctx.Next()
	}
}
