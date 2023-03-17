package httpapi

import (
	"strconv"

	"github.com/BetterWorks/gosk-api/internal/application"
	"github.com/BetterWorks/gosk-api/internal/types"
	"github.com/BetterWorks/gosk-api/internal/validation"
	"github.com/gofiber/fiber/v2"
)

// Config defines the input to NewServer
type Config struct {
	Application *application.Application `validate:"required"`
	BaseURL     string                   `validate:"required"`
	Logger      *types.Logger            `validate:"required"`
	Mode        string                   `validate:"required"`
	Namespace   string                   `validate:"required"`
	Port        uint                     `validate:"required"`
}

// Server defines a server for handling HTTP API requests
type Server struct {
	App         *fiber.App
	Logger      *types.Logger
	baseURL     string
	controllers *controllerRegistry
	namespace   string
	port        uint
}

// NewServer returns a new Server instance
func NewServer(c *Config) (*Server, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	app := fiber.New(fiber.Config{
		AppName:      c.Namespace,
		ErrorHandler: errorHandler,
	})

	log := c.Logger.Log.With().Str("tags", "httpapi").Logger()
	logger := &types.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	controllers := registerControllers(logger, c.Application.Services)

	s := &Server{
		Logger: logger,
		App:    app,
		// baseURL:    c.BaseURL,
		controllers: controllers,
		namespace:   c.Namespace,
		port:        c.Port,
	}

	s.configureMiddleware()
	s.registerRoutes()

	return s, nil
}

// Serve starts the HTTP server on the configured address
func (s *Server) Serve() error {
	addr := s.baseURL + ":" + strconv.FormatUint(uint64(s.port), 10)
	return s.App.Listen(addr)
}
