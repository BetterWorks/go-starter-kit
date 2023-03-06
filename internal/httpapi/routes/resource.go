package routes

import (
	"github.com/gofiber/fiber/v2"

	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
	"github.com/jasonsites/gosk-api/internal/types"
)

// ResourceRouter implements an example router group for a Resource resource
func ResourceRouter(r *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/resources"
	g := r.Group(prefix)

	// createResource provides a JSONRequestBody with data binding for the Resource model
	// for use with Create/Update Controller methods
	createResource := func() *types.JSONRequestBody {
		return &types.JSONRequestBody{
			Data: &types.RequestResource{
				Properties: &types.ResourceRequestData{},
			},
		}
	}

	g.Get("/", c.List())
	g.Get("/:id", c.Detail())
	g.Post("/", c.Create(createResource))
	g.Put("/:id", c.Update(createResource))
	g.Delete("/:id", c.Delete())
}
