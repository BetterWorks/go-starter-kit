package httpapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"

	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
	"github.com/jasonsites/gosk-api/internal/httpapi/routes"
)

// configureMiddleware
func (s *Server) configureMiddleware() {
	app := s.App

	skipHealth := func(ctx *fiber.Ctx) bool {
		return ctx.Path() == "/"+s.namespace+"/health"
	}

	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(mw.ResponseLogger(&mw.ResponseLoggerConfig{
		Logger: s.Logger,
		Next:   skipHealth,
	}))
	app.Use(helmet.New())
	app.Use(mw.Correlation(&mw.CorrelationConfig{
		Next: skipHealth,
	}))
	app.Use(mw.RequestLogger(&mw.RequestLoggerConfig{
		Logger: s.Logger,
		Next:   skipHealth,
	}))
}

// registerRoutes
func (s *Server) registerRoutes() {
	app := s.App
	c := s.controller
	ns := s.namespace

	routes.BaseRouter(app, c, ns)
	routes.HealthRouter(app, c, ns)

	routes.BookRouter(app, c, ns)
	routes.MovieRouter(app, c, ns)
}
