package domain

import (
	"github.com/BetterWorks/go-starter-kit/internal/core/app"
	"github.com/BetterWorks/go-starter-kit/internal/core/interfaces"
)

// Domain is the top-level container for the application domain layer
type Domain struct {
	Services *Services
}

// Services contains all individual domain services
type Services struct {
	Example interfaces.ExampleService `validate:"required"`
}

// NewDomain creates a new Domain instance
func NewDomain(s *Services) (*Domain, error) {
	if err := app.Validator.Validate.Struct(s); err != nil {
		return nil, err
	}

	return &Domain{Services: s}, nil
}
