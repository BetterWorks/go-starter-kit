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

// NewConflictError returns a new CustomError with the Conflict error type
func NewConflictError(err error, message string, a ...any) error {
	et := ErrorType.Conflict
	return wrapErrorf(err, et, message, a...)
}

// NewForbiddenError returns a new CustomError with the Forbidden error type
func NewForbiddenError(err error, message string, a ...any) error {
	et := ErrorType.Forbidden
	return wrapErrorf(err, et, message, a...)
}

// NewInternalServerError returns a new CustomError with the InternalServer error type
func NewInternalServerError(err error, message string, a ...any) error {
	et := ErrorType.InternalServer
	return wrapErrorf(err, et, message, a...)
}

// NewNotFoundError returns a new CustomError with the NotFound error type
func NewNotFoundError(err error, message string, a ...any) error {
	et := ErrorType.NotFound
	return wrapErrorf(err, et, message, a...)
}

// NewUnauthorizedError returns a new CustomError with the Unauthorized error type
func NewUnauthorizedError(err error, message string, a ...any) error {
	et := ErrorType.Unauthorized
	return wrapErrorf(err, et, message, a...)
}

// NewValidationError returns a new CustomError with the Validation error type
func NewValidationError(err error, message string, a ...any) error {
	et := ErrorType.Validation
	return wrapErrorf(err, et, message, a...)
}
