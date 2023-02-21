package routes

import (
	"github.com/gofiber/fiber/v2"

	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
	"github.com/jasonsites/gosk-api/internal/types"
)

// EpisodeRouter implements an example router group for an Episode resource
func EpisodeRouter(r *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/episodes"
	g := r.Group(prefix)

	// createResource provides a JSONRequestBody with data binding for the Episode model
	// for use with Create/Update Controller methods
	createResource := func() *types.JSONRequestBody {
		return &types.JSONRequestBody{
			Data: &types.RequestResource{
				Properties: &types.EpisodeRequestData{},
			},
		}
	}

	g.Get("/", c.List())
	g.Get("/:id", c.Detail())
	g.Post("/", c.Create(createResource))
	g.Put("/:id", c.Update(createResource))
	g.Delete("/:id", c.Delete())
}
