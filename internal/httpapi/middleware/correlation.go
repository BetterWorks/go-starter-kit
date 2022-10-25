package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CorrelationHeaders struct {
	XRequestID      string `header:"x-request-id"`
	XB3TraceID      string `header:"x-b3-traceid"`
	XB3SpanID       string `header:"x-b3-spanid"`
	XB3ParentSpanID string `header:"x-b3-parentspanid"`
	XB3Sampled      string `header:"x-b3-sampled"`
	XB3Flags        string `header:"x-b3-flags"`
	XOTSpanContext  string `header:"x-ot-span-context"`
}

func Correlation() gin.HandlerFunc {
	getTracingHeaders := func(ctx *gin.Context) *CorrelationHeaders {
		headers := &CorrelationHeaders{}

		if err := ctx.ShouldBindHeader(headers); err != nil {
			fmt.Printf("%v\n", headers) // @TODO
		}

		return headers
	}

	return func(ctx *gin.Context) {
		headers := getTracingHeaders(ctx)

		if headers.XRequestID == "" {
			headers.XRequestID = uuid.NewString()
		}

		ctx.Set("correlation", gin.H{
			"headers":   *headers,
			"requestId": headers.XRequestID,
		})
		ctx.Header("X-Request-ID", headers.XRequestID)
		ctx.Next()
	}
}
