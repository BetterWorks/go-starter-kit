package routes

import (
	"fmt"

	"github.com/BetterWorks/gosk-api/internal/core/models"
	ctrl "github.com/BetterWorks/gosk-api/internal/http/controllers"
	"github.com/go-chi/chi/v5"
)

// ExampleRouter implements a router group for an Example resource
func ExampleRouter(r *chi.Mux, c *ctrl.Controller, ns string) {
	prefix := fmt.Sprintf("/%s/examples", ns)

	// createResource provides a RequestBody with data binding for the Example model
	// for use with Create/Update Controller methods
	createResource := func() *ctrl.RequestBody {
		return &ctrl.RequestBody{
			Data: &ctrl.RequestResource{
				Attributes: &models.ExampleInputData{},
			},
		}
	}

	r.Route(prefix, func(r chi.Router) {
		// r.With(httpin.NewInput(ListUserReposInput{})).Get("/", c.List())
		r.Get("/", c.List())
		r.Get("/{id}", c.Detail())
		r.Post("/", c.Create(createResource))
		r.Put("/{id}", c.Update(createResource))
		r.Delete("/{id}", c.Delete())
	})
}
