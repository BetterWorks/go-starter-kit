package validation

import validator "github.com/go-playground/validator/v10"

// Validate provides an application validator
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}
