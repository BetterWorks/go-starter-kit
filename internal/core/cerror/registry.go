package cerror

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
