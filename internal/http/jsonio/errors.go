package jsonio

import (
	"errors"
	"fmt"
	"net/http"

	cerror "github.com/BetterWorks/go-starter-kit/internal/core/cerror"
	"github.com/BetterWorks/go-starter-kit/internal/core/jsonapi"
	"github.com/invopop/validation"
)

// HTTPStatusMap maps custom error types to relevant HTTP status codes
var HTTPStatusMap = map[string]int{
	cerror.ErrorType.Conflict:       http.StatusConflict,
	cerror.ErrorType.Forbidden:      http.StatusForbidden,
	cerror.ErrorType.InternalServer: http.StatusInternalServerError,
	cerror.ErrorType.NotFound:       http.StatusNotFound,
	cerror.ErrorType.Unauthorized:   http.StatusUnauthorized,
	cerror.ErrorType.Validation:     http.StatusBadRequest,
}

// EncodeError writes error messages to the response writer
func EncodeError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		code     = http.StatusInternalServerError
		response jsonapi.ErrorResponse
	)

	switch e := err.(type) {
	case cerror.CustomError:
		code, response = handleCustomError(e)
	default:
		data := defaultErrorData()
		response = errorResponse(data)
	}

	EncodeResponse(w, r, code, response)
}

func defaultErrorData() jsonapi.ErrorData {
	return jsonapi.ErrorData{
		Status: http.StatusInternalServerError,
		Title:  cerror.ErrorType.InternalServer,
		Detail: "internal server error",
	}
}

func defaultValidationErrorData(e cerror.CustomError) jsonapi.ErrorData {
	detail := "validation error"
	if e.ErrorMessage() != "" {
		detail = e.ErrorMessage()
	}

	return jsonapi.ErrorData{
		Status: http.StatusBadRequest,
		Title:  cerror.ErrorType.Validation,
		Detail: detail,
	}
}

func errorData(code int, detail, title string) jsonapi.ErrorData {
	return jsonapi.ErrorData{
		Status: code,
		Title:  title,
		Detail: detail,
	}
}

func errorResponse(data jsonapi.ErrorData) jsonapi.ErrorResponse {
	return jsonapi.ErrorResponse{
		Errors: []jsonapi.ErrorData{data},
	}
}

func handleCustomError(e cerror.CustomError) (int, jsonapi.ErrorResponse) {
	code := HTTPStatusMap[e.Type()]

	if e.Type() != cerror.ErrorType.Validation {
		data := errorData(code, e.ErrorMessage(), e.Type())
		return code, errorResponse(data)
	}

	var (
		response = errorResponse(defaultValidationErrorData(e))
		verrors  validation.Errors
	)
	if errors.As(e, &verrors) {
		response = validationErrorResponse("/", verrors, jsonapi.ErrorResponse{})
	}

	return code, response
}

func validationErrorResponse(path string, ve validation.Errors, er jsonapi.ErrorResponse) jsonapi.ErrorResponse {
	for key, val := range ve {
		path := fmt.Sprintf("%s%s", path, key)

		switch v := val.(type) {
		case validation.Error:
			er.Errors = append(er.Errors, jsonapi.ErrorData{
				Status: http.StatusBadRequest,
				Source: &jsonapi.ErrorSource{Pointer: path},
				Title:  cerror.ErrorType.Validation,
				Detail: v.Error(),
			})
		case validation.Errors:
			path = fmt.Sprintf("%s/", path)
			return validationErrorResponse(path, v, er)
		}
	}

	return er
}
