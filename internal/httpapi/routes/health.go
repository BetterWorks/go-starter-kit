package routes

import (
	"net/http"

	ctrl "github.com/BetterWorks/gosk-api/internal/httpapi/controllers"
	"github.com/gofiber/fiber/v2"
)

// HealthRouter implements an router group for healthcheck
func HealthRouter(r *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns + "/health"
	g := r.Group(prefix)

	status := func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{"status": "healthy"})
		return nil
	}

	g.Get("/", status)
}
