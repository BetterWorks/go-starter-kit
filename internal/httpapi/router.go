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
	ResourceController *controllers.Controller
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
	composeErrorResponse := func(code int, title, detail string) types.ErrorResponse {
		return types.ErrorResponse{
			Errors: []types.ErrorData{{
				Status: code,
				Title:  title,
				Detail: detail,
			}},
		}
	}

	var (
		code     = http.StatusInternalServerError // default error status code (500)
		detail   = "internal server error"
		title    string
		fiberErr *fiber.Error
		response types.ErrorResponse
	)

	cerr, ok := err.(*types.CustomError)
	// custom (controlled) errors
	if ok {
		code = types.HTTPStatusCodeMap[cerr.Type]
		detail = cerr.Message
		title = cerr.Type

		// fiber errors
	} else if errors.As(err, &fiberErr) {
		code = fiberErr.Code
		detail = fiberErr.Message

		// unknown errors
	} else {
		title = types.ErrorType.InternalServer
	}

	response = composeErrorResponse(code, title, detail)

	ctx.Status(code)
	if err := ctx.JSON(response); err != nil {
		return ctx.SendString(detail)
	}

	return nil
}

// registerControllers
func registerControllers(logger *types.Logger, services *application.Services) *controllerRegistry {
	return &controllerRegistry{
		ResourceController: controllers.NewController(&controllers.Config{
			Service: services.ResourceService,
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
	routes.ResourceRouter(app, c.ResourceController, ns)
}
