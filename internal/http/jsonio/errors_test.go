package jsonio

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/cerror"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
	tr "github.com/BetterWorks/go-starter-kit/test/testutils/reflection"
	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/invopop/validation"
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

func Test_EncodeError(t *testing.T) {
	tests := []struct {
		errFunc            func(error, string, ...any) error
		expectedStatusCode int
		expectedType       string
	}{{
		errFunc:            cerror.NewConflictError,
		expectedStatusCode: http.StatusConflict,
		expectedType:       cerror.ErrorType.Conflict,
	}, {
		errFunc:            cerror.NewForbiddenError,
		expectedStatusCode: http.StatusForbidden,
		expectedType:       cerror.ErrorType.Forbidden,
	}, {
		errFunc:            cerror.NewInternalServerError,
		expectedStatusCode: http.StatusInternalServerError,
		expectedType:       cerror.ErrorType.InternalServer,
	}, {
		errFunc:            cerror.NewNotFoundError,
		expectedStatusCode: http.StatusNotFound,
		expectedType:       cerror.ErrorType.NotFound,
	}, {
		errFunc:            cerror.NewUnauthorizedError,
		expectedStatusCode: http.StatusUnauthorized,
		expectedType:       cerror.ErrorType.Unauthorized,
	}}

	for _, tc := range tests {
		tc := tc
		name := fmt.Sprintf("Test EncodeError with %s ", tr.GetFunctionName(tc.errFunc))
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			errMessage := fake.Phrase()

			sourceErr := SourceError{}
			expectedErrorBody := fmt.Sprintf(`{"errors":[{"status":%d,"title":"%s","detail":"%s"}]}`, tc.expectedStatusCode, tc.expectedType, errMessage)

			err := tc.errFunc(sourceErr, errMessage)

			req := httptest.NewRequest("GET", "http://example.com/foo", nil)
			w := httptest.NewRecorder()

			EncodeError(w, req, err)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode != tc.expectedStatusCode {
				te.NewLineErrorf(t, tc.expectedStatusCode, resp.StatusCode)
			}

			strBody := strings.ReplaceAll(string(body), "\n", "")

			if strBody != expectedErrorBody {
				te.NewLineErrorf(t, expectedErrorBody, strBody)
			}
		})
	}
}

func Test_EncodeError_ValidationError(t *testing.T) {
	tests := []struct {
		errFunc            func(error, string, ...any) error
		expectedStatusCode int
		expectedType       string
		errorCount         int
	}{{
		errFunc:            cerror.NewValidationError,
		expectedStatusCode: http.StatusBadRequest,
		expectedType:       cerror.ErrorType.Validation,
		errorCount:         1,
	}, {
		errFunc:            cerror.NewValidationError,
		expectedStatusCode: http.StatusBadRequest,
		expectedType:       cerror.ErrorType.Validation,
		errorCount:         2,
	}}

	for _, tc := range tests {
		tc := tc
		name := fmt.Sprintf("Test EncodeError with %s ", tr.GetFunctionName(tc.errFunc))
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var sourceErr error
			errMessage := fake.Phrase()
			validationKeys := []string{}
			errors := validation.Errors{}
			errorStrings := []string{}

			for i := 0; i < tc.errorCount; i++ {
				validationKeys = append(validationKeys, fake.Word())
				errors[validationKeys[i]] = validation.NewError(tc.expectedType, errMessage)
				errorStrings = append(errorStrings, fmt.Sprintf(`{"status":%d,"source":{"pointer":"/%s"},"title":"%s","detail":"%s"}`, tc.expectedStatusCode, validationKeys[i], tc.expectedType, errMessage))
			}

			sourceErr = errors

			err := tc.errFunc(sourceErr, errMessage)

			req := httptest.NewRequest("GET", "http://example.com/foo", nil)
			w := httptest.NewRecorder()

			EncodeError(w, req, err)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode != tc.expectedStatusCode {
				te.NewLineErrorf(t, tc.expectedStatusCode, resp.StatusCode)
			}

			strBody := strings.ReplaceAll(string(body), "\n", "")

			if strings.Count(strBody, "status") != tc.errorCount {
				te.NewLineErrorf(t, tc.errorCount, strings.Count(strBody, "status"))
			}

			for _, errStr := range errorStrings {
				if !strings.Contains(strBody, errStr) {
					te.NewLineErrorf(t, errStr, strBody)
				}
			}
		})
	}
}
