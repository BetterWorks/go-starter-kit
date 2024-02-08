package jsonapi

// Envelope
type Envelope map[string]any

// RequestBody
type RequestBody struct {
	Data *RequestResource `json:"data" validate:"required"`
}

// RequestResource
type RequestResource struct {
	Type       string `json:"type" validate:"required"`
	ID         string `json:"id" validate:"omitempty,uuid4"`
	Attributes any    `json:"attributes" validate:"required"`
}
