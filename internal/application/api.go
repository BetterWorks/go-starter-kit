package application

import (
	"github.com/google/uuid"
	"github.com/jasonsites/gosk-api/internal/application/domain"
)

// Application
type Application struct {
	Services *Services
}

func NewApplication(s *Services) *Application {
	return &Application{Services: s}
}

// Services
type Services struct {
	EpisodeService Service
	SeasonService  Service
}

// Service
type Service interface {
	Create(any) (*domain.JSONResponseSingle, error)
	Delete(uuid.UUID) error
	Detail(uuid.UUID) (*domain.JSONResponseSingle, error)
	List(*domain.ListMeta) (*domain.JSONResponseMulti, error)
	Update(any) (*domain.JSONResponseSingle, error)
}
