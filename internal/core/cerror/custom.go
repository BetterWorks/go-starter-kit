package cerror

// CustomError
type CustomError struct {
	Message string
	Type    string
}

// Error
func (e *CustomError) Error() string {
	return e.Message
}

// NewConflictError
func NewConflictError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.Conflict,
	}
}

// NewForbiddenError
func NewForbiddenError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.Forbidden,
	}
}

// NewInternalServerError
func NewInternalServerError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.InternalServer,
	}
}

// NewNotFoundError
func NewNotFoundError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.NotFound,
	}
}

// NewUnauthorizedError
func NewUnauthorizedError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.Unauthorized,
	}
}

// NewValidationError
func NewValidationError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Type:    ErrorType.Validation,
	}
}
