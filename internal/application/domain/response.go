package domain

import (
	"github.com/google/uuid"
)

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
	ID         uuid.UUID         `json:"id"`
	Meta       *ResourceMetadata `json:"meta,omitempty"`
	Properties any               `json:"properties"` // TODO
}

// JSONResponseSingle
type JSONResponseSingle struct {
	Meta *APIMetadata      `json:"meta,omitempty"`
	Data *ResponseResource `json:"data"`
}

// Discover
func (r *JSONResponseSingle) Discover() Discoverable {
	return r
}

// JSONResponseMulti
type JSONResponseMulti struct {
	Meta *APIMetadata        `json:"meta"`
	Data *[]ResponseResource `json:"data"`
}

// Discover
func (r *JSONResponseMulti) Discover() Discoverable {
	return r
}
