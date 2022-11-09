package routes

import (
	"github.com/gofiber/fiber/v2"

	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// BookRouter implements an example router group for a Book domain resource
func BookRouter(app *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/books"
	g := app.Group(prefix)

	g.Get("/", c.BookList())
	g.Get("/:id", c.BookDetail())
	g.Post("/", c.BookCreate())
	g.Patch("/:id", c.BookUpdate())
	g.Delete("/:id", c.BookDelete())
}
