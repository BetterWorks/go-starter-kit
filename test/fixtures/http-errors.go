package fixtures

import (
	"net/http"

	"github.com/BetterWorks/go-starter-kit/internal/core/models"
)

// JSONAPIErrorDataBuilder
type JSONAPIErrorDataBuilder struct {
	code   int
	detail string
	title  string
	source *models.ErrorSource
}

func NewJSONAPIErrorDataBuilder() *JSONAPIErrorDataBuilder {
	return &JSONAPIErrorDataBuilder{
		code:   http.StatusInternalServerError,
		detail: "internal server error",
		title:  "InternalServerError",
		source: nil,
	}
}

func (b *JSONAPIErrorDataBuilder) Code(code int) *JSONAPIErrorDataBuilder {
	b.code = code
	return b
}

func (b *JSONAPIErrorDataBuilder) Detail(detail string) *JSONAPIErrorDataBuilder {
	b.detail = detail
	return b
}

func (b *JSONAPIErrorDataBuilder) Title(title string) *JSONAPIErrorDataBuilder {
	b.title = title
	return b
}

func (b *JSONAPIErrorDataBuilder) Source(source *models.ErrorSource) *JSONAPIErrorDataBuilder {
	b.source = source
	return b
}

func (b *JSONAPIErrorDataBuilder) Build() *models.ErrorData {
	return &models.ErrorData{
		Detail: b.detail,
		Status: b.code,
		Title:  b.title,
		Source: b.source,
	}
}

// JSONAPIErrorResponseBuilder
type JSONAPIErrorResponseBuilder struct {
	errors []models.ErrorData
}

func NewJSONAPIErrorResponseBuilder() *JSONAPIErrorResponseBuilder {
	return &JSONAPIErrorResponseBuilder{
		errors: []models.ErrorData{*NewJSONAPIErrorDataBuilder().Build()},
	}
}

func (b *JSONAPIErrorResponseBuilder) Errors(errors []models.ErrorData) *JSONAPIErrorResponseBuilder {
	b.errors = errors
	return b
}

func (b *JSONAPIErrorResponseBuilder) Build() *models.ErrorResponse {
	return &models.ErrorResponse{
		Errors: b.errors,
	}
}
