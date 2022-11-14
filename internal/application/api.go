package application

import "github.com/jasonsites/gosk-api/internal/application/domain"

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
	Create(any) (*domain.JSONResponseSolo, error)
	Delete(string) error
	Detail(string) (*domain.JSONResponseSolo, error)
	List(*domain.ListMeta) (*domain.JSONResponseMult, error)
	Update(any) (*domain.JSONResponseSolo, error)
}
