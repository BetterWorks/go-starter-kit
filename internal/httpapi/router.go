package httpapi

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"

	"github.com/jasonsites/gosk-api/internal/application"
	"github.com/jasonsites/gosk-api/internal/httpapi/controllers"
	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
	"github.com/jasonsites/gosk-api/internal/httpapi/routes"
	"github.com/jasonsites/gosk-api/internal/types"
)

type controllerRegistry struct {
	EpisodeController *controllers.Controller
	SeasonController  *controllers.Controller
}

// configureMiddleware
func (s *Server) configureMiddleware() {
	app := s.App

	skipHealth := func(ctx *fiber.Ctx) bool {
		return ctx.Path() == "/"+s.namespace+"/health"
	}

	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(mw.ResponseLogger(&mw.ResponseLoggerConfig{Logger: s.Logger, Next: skipHealth}))
	app.Use(helmet.New())
	app.Use(mw.Correlation(&mw.CorrelationConfig{Next: skipHealth}))
	app.Use(mw.RequestLogger(&mw.RequestLoggerConfig{Logger: s.Logger, Next: skipHealth}))
}

// errorHandler provides custom error handling (end of chain middleware) for all routes
func errorHandler(ctx *fiber.Ctx, err error) error {
	// default error status code (500)
	code := http.StatusInternalServerError

	// retrieve custom status code (and override default) if error is of type *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	response := types.CustomError{
		Errors: []types.ErrorData{{
			Status: code,
			Source: types.ErrorSource{Pointer: ""},
			Title:  "TodoErrorType",
			Detail: e.Message,
		}},
	}

	ctx.Status(code)
	if err := ctx.JSON(response); err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	return nil
}

// registerControllers
func registerControllers(logger *types.Logger, services *application.Services) *controllerRegistry {
	return &controllerRegistry{
		EpisodeController: controllers.NewController(&controllers.Config{
			Service: services.EpisodeService,
			Logger:  logger,
		}),
		SeasonController: controllers.NewController(&controllers.Config{
			Service: services.SeasonService,
			Logger:  logger,
		}),
	}
}

// registerRoutes
func (s *Server) registerRoutes() {
	app := s.App
	c := s.controllers
	ns := s.namespace

	routes.BaseRouter(app, nil, ns)
	routes.HealthRouter(app, nil, ns)

	routes.EpisodeRouter(app, c.EpisodeController, ns)
	routes.SeasonRouter(app, c.SeasonController, ns)
}
