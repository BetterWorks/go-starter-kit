package application

import (
	"github.com/jasonsites/gosk-api/internal/core/types"
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
	BookService  BookService
	MovieService MovieService
}

// BookService
type BookService interface {
	Create(*types.Book) (*types.JSONResponseDetail, error)
	Delete(id string) error
	Detail(id string) (*types.JSONResponseDetail, error)
	List(*types.ListMeta) (*types.JSONResponseList, error)
	Update(data *types.Book) (*types.JSONResponseDetail, error)
}

// MovieService
type MovieService interface {
	Create(*types.Movie) (*types.JSONResponseDetail, error)
	Delete(id string) error
	Detail(id string) (*types.JSONResponseDetail, error)
	List(*types.ListMeta) (*types.JSONResponseList, error)
	Update(data *types.Movie) (*types.JSONResponseDetail, error)
}
