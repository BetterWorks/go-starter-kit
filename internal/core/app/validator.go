package app

import "github.com/go-playground/validator/v10"

type appValidator struct {
	Validate *validator.Validate
}

var Validator = appValidator{
	Validate: validator.New(),
}
