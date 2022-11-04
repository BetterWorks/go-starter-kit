package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/core/types"
)

// TODO https://opentracing.io/
// Correlation
func Correlation() gin.HandlerFunc {
	getTracingHeaders := func(ctx *gin.Context) *types.TracingHeaders {
		headers := &types.TracingHeaders{}

		if err := ctx.ShouldBindHeader(headers); err != nil {
			fmt.Printf("error binding `headers`: %+v\n", headers) // TODO
		}

		return headers
	}

	return func(ctx *gin.Context) {
		headers := getTracingHeaders(ctx)
		if headers.XRequestID == "" {
			headers.XRequestID = uuid.NewString()
		}

		trace := types.Trace{
			Headers:   *headers,
			RequestID: headers.XRequestID,
		}

		ctx.Set("Trace", trace)
		ctx.Header("x-request-id", trace.RequestID)
		ctx.Next()
	}
}
