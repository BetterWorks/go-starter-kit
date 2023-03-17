package routes

import (
	"net/http"

	ctrl "github.com/BetterWorks/gosk-api/internal/httpapi/controllers"
	"github.com/gofiber/fiber/v2"
)

// BaseRouter only exists to easily verify a working app and should normally be removed
func BaseRouter(r *fiber.App, c *ctrl.Controller, ns string) {
	prefix := "/" + ns
	g := r.Group(prefix)

	get := func(ctx *fiber.Ctx) error {
		headers := ctx.GetReqHeaders()
		host := ctx.Hostname()
		path := ctx.Path()
		remoteAddress := ctx.Context().RemoteAddr()

		ctx.Status(http.StatusOK)
		ctx.JSON(fiber.Map{
			"data": "base router is working...",
			"request": fiber.Map{
				"headers":       headers,
				"host":          host,
				"path":          path,
				"remoteAddress": remoteAddress,
			},
		})

		return nil
	}

	g.Get("/", get)
}
