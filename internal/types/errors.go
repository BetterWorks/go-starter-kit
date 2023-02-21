package types

const ValidationErrorType = "ValidationError"

// ErrorResponse
type CustomError struct {
	Type   string      `json:"-"`
	Errors []ErrorData `json:"errors"`
}

// ValidationError
type ErrorData struct {
	Status int         `json:"status"`
	Source ErrorSource `json:"source"`
	Title  string      `json:"title"`
	Detail string      `json:"detail"`
}

// ErrorSource
type ErrorSource struct {
	Pointer string `json:"pointer"`
}
