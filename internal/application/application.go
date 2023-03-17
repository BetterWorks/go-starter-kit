package application

import (
	"github.com/BetterWorks/gosk-api/internal/types"
)

// Application
type Application struct {
	Services *Services
}

// Services
type Services struct {
	ResourceService types.Service
}

func NewApplication(s *Services) *Application {
	return &Application{Services: s}
}
