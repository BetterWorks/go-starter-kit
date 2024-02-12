package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/BetterWorks/go-starter-kit/internal/core/app"
	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/domain"
	ctrl "github.com/BetterWorks/go-starter-kit/internal/http/controllers"
	"github.com/go-chi/chi/v5"
)

// ServerConfig defines the input to NewServer
type ServerConfig struct {
	Domain       *domain.Domain       `validate:"required"`
	Host         string               `validate:"required"`
	Logger       *logger.CustomLogger `validate:"required"`
	Port         uint                 `validate:"required"`
	QueryConfig  *ctrl.QueryConfig    `validate:"required"`
	RouterConfig *RouterConfig        `validate:"required"`
}

// Server defines a server for handling HTTP API requests
type Server struct {
	Logger *logger.CustomLogger
	Port   uint
	Server *http.Server
}

// NewServer returns a new Server instance
func NewServer(c *ServerConfig) (*Server, error) {
	if err := app.Validator.Validate.Struct(c); err != nil {
		return nil, err
	}

	mux := chi.NewRouter()
	controllers, err := registerControllers(c.Domain.Services, c.Logger, c.QueryConfig)
	if err != nil {
		return nil, err
	}

	configureMiddleware(c.RouterConfig, mux, c.Logger)
	registerRoutes(c.RouterConfig, mux, controllers)

	addr := fmt.Sprintf(":%s", strconv.FormatUint(uint64(c.Port), 10))
	s := &Server{
		Logger: c.Logger,
		Port:   c.Port,
		Server: &http.Server{Addr: addr, Handler: mux},
	}

	return s, nil
}

// Serve starts the HTTP server on the configured address
func (s *Server) Serve() error {
	s.Logger.Log.Info(fmt.Sprintf("server listening on port :%d", s.Port))
	return s.Server.ListenAndServe()
}
