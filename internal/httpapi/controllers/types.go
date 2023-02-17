package controllers

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

const ValidationErrorType = "ValidationError"

// ValidationErrorSource
type ValidationErrorSource struct {
	Pointer string `json:"pointer"`
}

// ValidationError
type ValidationError struct {
	Status int                   `json:"status"`
	Source ValidationErrorSource `json:"source"`
	Title  string                `json:"title"`
	Detail string                `json:"detail"`
}

// ErrorResponse
type ErrorResponse struct {
	Errors []ValidationError `json:"errors"`
}
