package types

// JSON API Response ------------------------------------------------------------------------------
// JSONResponse defines the interface for a JSON REST response
type JSONResponse interface {
	Discoverable
}

// JSONResponseDetail
type JSONResponseDetail struct {
	Meta *APIMetadata      `json:"meta,omitempty"`
	Data *ResponseResource `json:"data"`
}

// Discover
func (r *JSONResponseDetail) Discover() Discoverable {
	return r
}

// JSONResponseList
type JSONResponseList struct {
	Meta *APIMetadata        `json:"meta"`
	Data *[]ResponseResource `json:"data"`
}

// Discover
func (r *JSONResponseList) Discover() Discoverable {
	return r
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
