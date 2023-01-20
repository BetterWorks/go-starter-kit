package application

import (
	"github.com/jasonsites/gosk-api/internal/types"
)

// Application
type Application struct {
	Services *Services
}

// Services
type Services struct {
	EpisodeService types.Service
	SeasonService  types.Service
}

func NewApplication(s *Services) *Application {
	return &Application{Services: s}
}
