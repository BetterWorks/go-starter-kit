package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jasonsites/gosk-api/internal/application/domain"
	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// SeasonRouter implements an example router group for a Season resource
func SeasonRouter(app *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/seasons"
	g := app.Group(prefix)

	// HandlerFunc wrapper that injects Season for request body data binding
	wrapper := func(f func(*ctrl.JSONRequestBody) fiber.Handler) fiber.Handler {
		resource := &ctrl.JSONRequestBody{Data: &ctrl.RequestResource{Properties: &domain.Season{}}}
		return f(resource)
	}

	g.Get("/", c.List())
	g.Get("/:id", c.Detail())
	g.Post("/", wrapper(c.Create))
	g.Put("/:id", wrapper(c.Update))
	g.Delete("/:id", c.Delete())
}
