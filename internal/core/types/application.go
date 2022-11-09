package types

import "github.com/rs/zerolog"

// Context ----------------------------------------------------------------------------------------
// Trace
type Trace struct {
	Headers   map[string]string
	RequestID string
}

// List Metadata ----------------------------------------------------------------------------------
// ListFilters
type ListFilters struct{}

// ListMeta
type ListMeta struct {
	Filters ListFilters
	Paging  ListPaging
	Sorting ListSorting
}

// ListPaging
type ListPaging struct {
	Limit  uint
	Offset uint
	Total  uint
}

// ListSorting
type ListSorting struct {
}

// List Metadata ----------------------------------------------------------------------------------
// Logger encapsulates a logger with an associated log level and toggle
type Logger struct {
	Enabled bool
	Level   string
	Log     *zerolog.Logger
}
