package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jasonsites/gosk-api/internal/core/types"
	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// BookRouter implements an example router group for a Book domain resource
func BookRouter(app *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/books"
	t := types.ResourceType.Book
	g := app.Group(prefix)

	// HandlerFunc wrapper that injects Book for request body data binding
	wrapper := func(f func(any, string) fiber.Handler) fiber.Handler {
		return f(&types.Book{}, t)
	}

	g.Get("/", c.List(t))
	g.Get("/:id", c.Detail(t))
	g.Post("/", wrapper(c.Create))
	g.Patch("/:id", wrapper(c.Update))
	g.Delete("/:id", c.Delete(t))
}
