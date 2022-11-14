package domain

// JSONResponse defines the interface for a JSON serialized response
type JSONResponse interface {
	Discoverable
}

// APIMetadata
type APIMetadata struct {
	Paging *ListPaging `json:"paging,omitempty"`
}

// ResourceMetadata
type ResourceMetadata struct{}

// Resource
type ResponseResource struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Meta       *ResourceMetadata `json:"meta,omitempty"`
	Properties any               `json:"properties"` // TODO
}

// JSONResponseSolo
type JSONResponseSolo struct {
	Meta *APIMetadata      `json:"meta,omitempty"`
	Data *ResponseResource `json:"data"`
}

// Discover
func (r *JSONResponseSolo) Discover() Discoverable {
	return r
}

// JSONResponseMult
type JSONResponseMult struct {
	Meta *APIMetadata        `json:"meta"`
	Data *[]ResponseResource `json:"data"`
}

// Discover
func (r *JSONResponseMult) Discover() Discoverable {
	return r
}
