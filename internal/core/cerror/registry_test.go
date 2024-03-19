package cerror

import (
	"fmt"
	"testing"

	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	tr "github.com/BetterWorks/go-starter-kit/test/testutils/reflection"
	fake "github.com/brianvoe/gofakeit/v6"
)

func TestErrorRegistry(t *testing.T) {
	sourceErr := SourceError{}

	errMessage := fake.Phrase()
	expectedError := fmt.Sprintf("%s: source error", errMessage)
	expectedErrorMessage := errMessage

	tests := []struct {
		errFunc      func(error, string, ...any) error
		expectedType string
	}{{
		errFunc:      NewConflictError,
		expectedType: ErrorType.Conflict,
	}, {
		errFunc:      NewForbiddenError,
		expectedType: ErrorType.Forbidden,
	}, {
		errFunc:      NewInternalServerError,
		expectedType: ErrorType.InternalServer,
	}, {
		errFunc:      NewNotFoundError,
		expectedType: ErrorType.NotFound,
	}, {
		errFunc:      NewUnauthorizedError,
		expectedType: ErrorType.Unauthorized,
	}, {
		errFunc:      NewValidationError,
		expectedType: ErrorType.Validation,
	}}

	for _, tc := range tests {
		tc := tc
		name := fmt.Sprintf("Test %s", tr.GetFunctionName(tc.errFunc))
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := tc.errFunc(sourceErr, errMessage)

			if err.Error() != expectedError {
				te.NewLineErrorf(t, expectedError, err.Error())
			}

			customErr := err.(CustomError)

			if customErr.ErrorMessage() != expectedErrorMessage {
				te.NewLineErrorf(t, expectedErrorMessage, customErr.ErrorMessage())
			}

			if customErr.Type() != tc.expectedType {
				te.NewLineErrorf(t, tc.expectedType, customErr.Type())
			}

			if customErr.Unwrap() != sourceErr {
				te.NewLineErrorf(t, sourceErr, customErr.Unwrap())
			}
		})
	}
}
