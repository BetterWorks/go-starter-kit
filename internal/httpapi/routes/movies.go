package routes

import (
	"github.com/gofiber/fiber/v2"

	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// MovieRouter implements an example router group for a Movie resource
func MovieRouter(r *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/movies"
	g := r.Group(prefix)

	g.Get("/", c.MovieList())
	g.Get("/:id", c.MovieDetail())
	g.Post("/", c.MovieCreate())
	g.Patch("/:id", c.MovieUpdate())
	g.Delete("/:id", c.MovieDelete())
}
