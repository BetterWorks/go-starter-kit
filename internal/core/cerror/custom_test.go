package cerror

import (
	"fmt"
	"testing"

	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	fake "github.com/brianvoe/gofakeit/v6"
)

type SourceError struct {
	message string
}

func (e SourceError) Error() string {
	if e.message != "" {
		return e.message
	}
	return "source error"
}

func TestWrapErrorf(t *testing.T) {
	sourceErr := SourceError{}

	errType := fake.Word()
	errMessage := fake.Phrase()

	tests := []struct {
		name                 string
		errType              string
		message              string
		expectedError        string
		expectedErrorMessage string
		expectedType         string
	}{{
		name:                 "Test providing message",
		errType:              errType,
		message:              errMessage,
		expectedError:        fmt.Sprintf("%s: source error", errMessage),
		expectedErrorMessage: errMessage,
		expectedType:         errType,
	}, {
		name:                 "Test not providing message",
		errType:              errType,
		message:              "",
		expectedError:        fmt.Sprintf("%s: source error", errType),
		expectedErrorMessage: errType,
		expectedType:         errType,
	}}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := wrapErrorf(sourceErr, tc.errType, tc.message)

			if err.Error() != tc.expectedError {
				te.NewLineErrorf(t, tc.expectedError, err.Error())
			}

			customErr := err.(CustomError)

			if customErr.ErrorMessage() != tc.expectedErrorMessage {
				te.NewLineErrorf(t, tc.expectedErrorMessage, customErr.ErrorMessage())
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
