package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jasonsites/gosk-api/internal/core/types"
	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// MovieRouter implements an example router group for a Movie domain resource
func MovieRouter(r *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/movies"
	t := types.ResourceType.Movie
	g := r.Group(prefix)

	// HandlerFunc wrapper that injects Movie for request body data binding
	wrapper := func(f func(any, string) fiber.Handler) fiber.Handler {
		return f(&types.Movie{}, t)
	}

	g.Get("/", c.List(t))
	g.Get("/:id", c.Detail(t))
	g.Post("/", wrapper(c.Create))
	g.Patch("/:id", wrapper(c.Update))
	g.Delete("/:id", c.Delete(t))
}
