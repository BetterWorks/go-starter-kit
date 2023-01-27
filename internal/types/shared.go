package types

import (
	"encoding/json"

	"github.com/rs/zerolog"
)

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

// Nullable Types ---------------------------------------------------------------------------------
// NullBool
type NullBool struct {
	Valid bool
	Value bool
}

func (n *NullBool) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(n.Value)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (n *NullBool) UnmarshalJSON(b []byte) error {
	var unmarshalledJson bool

	err := json.Unmarshal(b, &unmarshalledJson)
	if err != nil {
		return err
	}

	n.Value = unmarshalledJson
	n.Valid = true

	return nil
}

// NullUint32
type NullUint32 struct {
	Uint32 uint32
	Valid  bool
}
