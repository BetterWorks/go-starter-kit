package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Trace
type Trace struct {
	Headers   TracingHeaders
	RequestID string
}

// TracingHeaders
type TracingHeaders struct {
	XRequestID      string `header:"x-request-id"`
	XB3TraceID      string `header:"x-b3-traceid"`
	XB3SpanID       string `header:"x-b3-spanid"`
	XB3ParentSpanID string `header:"x-b3-parentspanid"`
	XB3Sampled      string `header:"x-b3-sampled"`
	XB3Flags        string `header:"x-b3-flags"`
	XOTSpanContext  string `header:"x-ot-span-context"`
}

// TODO https://opentracing.io/
// Correlation
func Correlation() gin.HandlerFunc {
	getTracingHeaders := func(ctx *gin.Context) *TracingHeaders {
		headers := &TracingHeaders{}

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

		trace := Trace{
			Headers:   *headers,
			RequestID: headers.XRequestID,
		}

		ctx.Set("Trace", trace)
		ctx.Header("x-request-id", trace.RequestID)
		ctx.Next()
	}
}
