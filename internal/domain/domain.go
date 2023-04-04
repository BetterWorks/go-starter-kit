package domain

import (
	"github.com/BetterWorks/gosk-api/internal/types"
)

// Domain is the top-level container for the application domain layer
type Domain struct {
	Services *Services
}

// Services contains all individual resource services
type Services struct {
	ResourceService types.Service
}

func NewDomain(s *Services) *Domain {
	return &Domain{Services: s}
}
