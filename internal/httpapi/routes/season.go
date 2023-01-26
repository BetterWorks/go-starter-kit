package routes

import (
	"github.com/gofiber/fiber/v2"

	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
	"github.com/jasonsites/gosk-api/internal/types"
)

// SeasonRouter implements an example router group for a Season resource
func SeasonRouter(r *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/seasons"
	g := r.Group(prefix)

	// createResource provides a JSONRequestBody with data binding for the Season model
	// for use with Create/Update Controller methods
	createResource := func() *ctrl.JSONRequestBody {
		return &ctrl.JSONRequestBody{
			Data: &ctrl.RequestResource{
				Properties: &types.SeasonRequestData{},
			},
		}
	}

	g.Get("/", c.List())
	g.Get("/:id", c.Detail())
	g.Post("/", c.Create(createResource))
	g.Put("/:id", c.Update(createResource))
	g.Delete("/:id", c.Delete())
}
