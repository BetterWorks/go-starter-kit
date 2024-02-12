package cerror

import "fmt"

// CustomError
type CustomError struct {
	errorType string
	message   string
	sourceErr error
}

// Error returns the custom error message with the source error message (if present)
func (e CustomError) Error() string {
	if e.sourceErr != nil {
		return fmt.Sprintf("%s: %v", e.message, e.sourceErr)
	}

	return e.message
}

// ErrorMessage returns the custom error message
func (e CustomError) ErrorMessage() string {
	return e.message
}

// Type returns the error type string constant
func (e CustomError) Type() string {
	return e.errorType
}

// Unwrap returns the wrapped source error
func (e CustomError) Unwrap() error {
	return e.sourceErr
}

// wrapErrorf
func wrapErrorf(err error, errtype, message string, a ...any) error {
	if message == "" {
		message = errtype
	}
	return CustomError{
		errorType: errtype,
		message:   fmt.Sprintf(message, a...),
		sourceErr: err,
	}
}
