package types

import "github.com/rs/zerolog"

// Context ----------------------------------------------------------------------------------------

// CorrelationContextKey
const CorrelationContextKey string = "Trace"

// Logger encapsulates a logger with an associated log level and toggle
type Logger struct {
	Enabled bool
	Level   string
	Log     *zerolog.Logger
}

// Trace
type Trace struct {
	Headers   map[string]string
	RequestID string // TODO: consider uuid.UUID
}

// JSON Request -----------------------------------------------------------------------------------

// JSONRequestBody
type JSONRequestBody struct {
	Data *RequestResource `json:"data" validate:"required"`
}

// RequestResource
type RequestResource struct {
	Type       string `json:"type" validate:"required"`
	ID         string `json:"id" validate:"omitempty,uuid4"`
	Properties any    `json:"properties" validate:"required"`
}

// Query ------------------------------------------------------------------------------------------

// QueryData composes all query parameters into a single struct for use across the app
type QueryData struct {
	Filters QueryFilters `query:"f"`
	Options QueryOptions `query:"o"`
	Paging  QueryPaging  `query:"p"`
	Sorting QuerySorting `query:"s"`
}

// QueryFilters defines the filter-related query paramaters
// f[enabled]=true&f[name]=test&f[status]=4
type QueryFilters struct {
	Enabled *bool   `query:"enabled"`
	Name    *string `query:"name"`
	Status  *int    `query:"status"`
}

// QueryOptions defines the options-related query paramaters
// o[export]=true
type QueryOptions struct {
	Export *bool `query:"export"`
}

// QueryPaging defines the paging-related query paramaters
// p[limit]=20&p[offset]=10
type QueryPaging struct {
	Limit  *int `query:"limit"`
	Offset *int `query:"offset"`
}

// QuerySorting defines the sorting-related query paramaters
// s[order]=desc&s[prop]=name
type QuerySorting struct {
	Order *string `query:"order"`
	Prop  *string `query:"prop"`
}
