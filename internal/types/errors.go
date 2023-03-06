package types

import "net/http"

// HTTP Response Errors ---------------------------------------------------------------------------

// ErrorResponse
type ErrorResponse struct {
	Errors []ErrorData `json:"errors"`
}

// ValidationError
type ErrorData struct {
	Status int          `json:"status"`
	Source *ErrorSource `json:"source,omitempty"`
	Title  string       `json:"title"`
	Detail string       `json:"detail"`
}

// ErrorSource
type ErrorSource struct {
	Pointer string `json:"pointer,omitempty"`
}

var HTTPStatusCodeMap = map[string]int{
	ErrorType.Conflict:       http.StatusConflict,
	ErrorType.Forbidden:      http.StatusForbidden,
	ErrorType.InternalServer: http.StatusInternalServerError,
	ErrorType.NotFound:       http.StatusNotFound,
	ErrorType.Unauthorized:   http.StatusUnauthorized,
	ErrorType.Validation:     http.StatusBadRequest,
}

// ------------------------------------------------------------------------------------------------

// ErrorRegistry defines a registry for all errors to be used across the application
type ErrorRegistry struct {
	Conflict       string
	Forbidden      string
	InternalServer string
	NotFound       string
	Unauthorized   string
	Validation     string
}

// ErrorType exposes constants for all error types
var ErrorType = ErrorRegistry{
	Conflict:       "ConflictError",
	Forbidden:      "ForbiddenError",
	InternalServer: "InternalServerError",
	NotFound:       "NotFoundError",
	Unauthorized:   "UnauthorizedError",
	Validation:     "ValidationError",
}

// ------------------------------------------------------------------------------------------------

// CustomError
type CustomError struct {
	Message string
	Type    string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewConflictError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.Conflict,
	}
}

func NewForbiddenError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.Forbidden,
	}
}

func NewInternalServerError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.InternalServer,
	}
}

func NewNotFoundError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.NotFound,
	}
}

func NewUnauthorizedError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.Unauthorized,
	}
}

func NewValidationError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.Validation,
	}
}
