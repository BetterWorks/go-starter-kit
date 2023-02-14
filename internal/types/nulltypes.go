package types

import (
	"encoding/json"
)

// Boolean ----------------------------------------------------------------------------------------

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

	n.Valid = true
	n.Value = unmarshalledJson

	return nil
}

// Integer ----------------------------------------------------------------------------------------

// NullInt
type NullInt struct {
	Valid bool
	Value int
}

// String -----------------------------------------------------------------------------------------

// NullString
type NullString struct {
	Valid bool
	Value string
}

// Unsigned 32bit Integer -------------------------------------------------------------------------

// NullUint32
type NullUint32 struct {
	Valid bool
	Value uint32
}
