package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	ctrl "github.com/jasonsites/gosk-api/internal/httpapi/controllers"
)

// HealthRouter implements an router group for healthcheck
func HealthRouter(app *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/health"
	g := app.Group(prefix)

	status := func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{"status": "healthy"})
		return nil
	}

	g.Get("/", status)
}
