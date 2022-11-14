package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/jasonsites/gosk-api/internal/application/domain"
	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// EpisodeRouter implements an example router group for an Episode resource
func EpisodeRouter(r *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/episodes"
	g := r.Group(prefix)
	fmt.Println(g)

	// HandlerFunc wrapper that injects Episode for request body data binding
	wrapper := func(f func(*ctrl.JSONRequestBody) fiber.Handler) fiber.Handler {
		resource := &ctrl.JSONRequestBody{Data: &ctrl.RequestResource{Properties: &domain.Episode{}}}
		return f(resource)
	}

	g.Get("/", c.List())
	g.Get("/:id", c.Detail())
	g.Post("/", wrapper(c.Create))
	g.Patch("/:id", wrapper(c.Update))
	g.Delete("/:id", c.Delete())
}
