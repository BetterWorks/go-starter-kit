package httpapi

import (
	"strconv"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/jasonsites/gosk-api/internal/core"
	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
	"github.com/jasonsites/gosk-api/internal/httpapi/routes"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// Config defines the input to NewServer
type Config struct {
	BaseURL   string       `validate:"required"`
	Logger    *core.Logger `validate:"required"`
	Namespace string       `validate:"required"`
	Port      uint         `validate:"required"`
}

// Server defines a server for handling HTTP API requests
type Server struct {
	Logger     *core.Logger
	Router     *gin.Engine
	baseURL    string
	controller *Controller
	namespace  string
	port       uint
}

// NewServer returns a new Server instance
func NewServer(c *Config) (*Server, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	r := gin.New()
	ctrl := newController()

	log := c.Logger.Log.With().Str("label", "httpapi").Logger()

	logger := &core.Logger{
		Enabled: c.Logger.Enabled,
		Level:   c.Logger.Level,
		Log:     &log,
	}

	s := &Server{
		Logger:     logger,
		Router:     r,
		baseURL:    c.BaseURL,
		controller: ctrl,
		namespace:  c.Namespace,
		port:       c.Port,
	}

	s.configureMiddleware()
	s.registerRoutes()

	return s, nil
}

func (s *Server) Serve() {
	addr := ":" + strconv.FormatUint(uint64(s.port), 10)
	s.Router.Run(addr)
}

func (s *Server) configureMiddleware() {
	r := s.Router

	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(mw.ResponseLogger(s.Logger))
	r.Use(mw.ResponseTime())
	// security headers
	// "github.com/gin-contrib/cors"
	r.Use(mw.ErrorHandler())
	r.Use(mw.Correlation())
	r.Use(mw.RequestLogger(s.Logger))
}

func (s *Server) registerRoutes() {
	c := s.controller
	ns := s.namespace
	r := s.Router

	routes.Health(nil, ns, r)
	routes.Base(nil, ns, r)
	routes.Resource(c, ns, r)
}
