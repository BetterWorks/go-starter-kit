package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jasonsites/gosk-api/internal/validation"
	"github.com/rs/zerolog"
)

// validateBody validates tagged fields in json request body
func validateBody(body *JSONRequestBody, log zerolog.Logger) *ErrorResponse {
	var errors []ValidationError

	if err := validation.Validate.Struct(body); err != nil {
		log.Error().Err(err).Msg("")

		for _, err := range err.(validator.ValidationErrors) {

			var (
				field   = strings.ToLower(err.Field())
				param   = err.Param()
				pointer = formatPath(err.Namespace())
				tag     = strings.ToLower(err.Tag())
			)

			// TODO: validation error field reference
			// fmt.Printf("\n\nNamespace: %s\n", err.Namespace())
			// fmt.Printf("Field: %s\n", err.Field())
			// fmt.Printf("StructNamespace: %s\n", err.StructNamespace())
			// fmt.Printf("StructField: %s\n", err.StructField())
			// fmt.Printf("Tag: %s\n", err.Tag())
			// fmt.Printf("ActualTag: %s\n", err.ActualTag())
			// fmt.Printf("Kind: %+v\n", err.Kind())
			// fmt.Printf("Type: %+v\n", err.Type())
			// fmt.Printf("Value: %+v\n", err.Value())
			// fmt.Printf("Param: %s\n\n", err.Param())

			verr := ValidationError{
				Status: http.StatusBadRequest,
				Source: ValidationErrorSource{Pointer: pointer},
				Title:  ValidationErrorType,
				Detail: formatErrorDetail(field, param, tag),
			}

			errors = append(errors, verr)
		}

		return &ErrorResponse{Errors: errors}
	}

	return nil
}

// formatErrorDetail builds the error message detail from the validation error's field/tag data
func formatErrorDetail(field, param, tag string) string {
	switch tag {
	// TODO: other relevant validation cases
	case "max":
		return fmt.Sprintf("'%s' field must contain a maximum of %s characters", field, param)
	case "min":
		return fmt.Sprintf("'%s' field must contain at least %s characters", field, param)
	case "required":
		return fmt.Sprintf("'%s' field is %s", field, tag)
	default:
		return fmt.Sprintf("validation error on field '%s' with tag '%s'", field, tag)
	}

}

// formatPath builds the error message's source.pointer path from the validation error's namespace
func formatPath(ns string) string {
	var b bytes.Buffer

	segments := strings.Split(ns, ".")
	for _, seg := range segments {
		if seg != "JSONRequestBody" {
			b.WriteString("/")
			b.WriteString(strings.ToLower(seg))
		}
	}

	return b.String()
}
