package controllers

import (
	"net/http"

	cerror "github.com/BetterWorks/gosk-api/internal/core/cerror"
	"github.com/BetterWorks/gosk-api/internal/core/jsonapi"
)

// HTTPStatusCodeMap maps custom error types to relevant HTTP status codes
var HTTPStatusCodeMap = map[string]int{
	cerror.ErrorType.Conflict:       http.StatusConflict,
	cerror.ErrorType.Forbidden:      http.StatusForbidden,
	cerror.ErrorType.InternalServer: http.StatusInternalServerError,
	cerror.ErrorType.NotFound:       http.StatusNotFound,
	cerror.ErrorType.Unauthorized:   http.StatusUnauthorized,
	cerror.ErrorType.Validation:     http.StatusBadRequest,
}

func (c *Controller) Error(w http.ResponseWriter, r *http.Request, err error) {
	var (
		code     = http.StatusInternalServerError
		detail   = "internal server error"
		title    string
		response jsonapi.ErrorResponse
	)

	// span := trace.SpanFromContext(r.Context())
	// span.RecordError(err)

	switch e := err.(type) {
	case *cerror.CustomError:
		code = HTTPStatusCodeMap[e.Type]
		title = e.Type
		if e.Type != cerror.ErrorType.InternalServer {
			detail = e.Message
		}
	default:
		title = cerror.ErrorType.InternalServer
	}

	response = composeErrorResponse(code, title, detail)
	c.JSONEncode(w, r, code, response)
}

func composeErrorResponse(code int, title, detail string) jsonapi.ErrorResponse {
	return jsonapi.ErrorResponse{
		Errors: []jsonapi.ErrorData{{
			Status: code,
			Title:  title,
			Detail: detail,
		}},
	}
}
