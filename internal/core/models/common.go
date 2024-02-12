package models

import "github.com/BetterWorks/go-starter-kit/internal/core/pagination"

// Envelope
type Envelope map[string]any

// Response
type Response struct {
	Meta *ResponseMetadata `json:"meta"`
	Data any               `json:"data"`
}

// ResponseMetadata
type ResponseMetadata struct {
	Page pagination.PageMetadata `json:"page,omitempty"`
}
