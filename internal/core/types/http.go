package types

// REQUEST ----------------------------------------------------------------------------------------
// RequestFilters
type RequestFilters struct{}

// RequestMeta
type RequestMeta struct {
	Filters RequestFilters
	Paging  RequestPaging
	Sorting RequestSorting
}

// RequestPaging
type RequestPaging struct{}

// RequestSorting
type RequestSorting struct{}

// Trace
type Trace struct {
	Headers   map[string]string
	RequestID string
}

// // TracingHeaders
// type TracingHeaders struct {
// 	XRequestID      string `header:"x-request-id"`
// 	XB3TraceID      string `header:"x-b3-traceid"`
// 	XB3SpanID       string `header:"x-b3-spanid"`
// 	XB3ParentSpanID string `header:"x-b3-parentspanid"`
// 	XB3Sampled      string `header:"x-b3-sampled"`
// 	XB3Flags        string `header:"x-b3-flags"`
// 	XOTSpanContext  string `header:"x-ot-span-context"`
// }
