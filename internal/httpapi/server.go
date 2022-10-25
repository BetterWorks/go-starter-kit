package httpapi

import (
	"strconv"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
	"github.com/jasonsites/gosk-api/internal/httpapi/routes"
	"github.com/jasonsites/gosk-api/internal/validation"
	"github.com/sirupsen/logrus"
)

// Config defines the input to NewServer
type Config struct {
	BaseURL   string             `validate:"required"`
	Log       logrus.FieldLogger `validate:"required"`
	Namespace string             `validate:"required"`
	Port      uint               `validate:"required"`
}

// Server defines a server for handling HTTP API requests
type Server struct {
	Router     *gin.Engine
	baseURL    string
	controller *Controller
	log        logrus.FieldLogger
	namespace  string
	port       uint
}

// NewServer returns a new Server instance
func NewServer(c *Config) (*Server, error) {
	if err := validation.Validate.Struct(c); err != nil {
		return nil, err
	}

	r := gin.Default()
	ctrl := newController()

	s := &Server{
		Router:     r,
		baseURL:    c.BaseURL,
		controller: ctrl,
		log:        c.Log, // @TODO not yet sure how to integrate app logger with gin's builtin logger
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

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(mw.ResponseLogger())
	r.Use(mw.ResponseTime())
	// security headers
	// "github.com/gin-contrib/cors"
	r.Use(mw.ErrorHandler())
	r.Use(mw.Correlation())
	r.Use(mw.RequestLogger())
}

func (s *Server) registerRoutes() {
	c := s.controller
	ns := s.namespace
	r := s.Router

	routes.Health(nil, ns, r)
	routes.Base(nil, ns, r)
	routes.Resource(c, ns, r)
}
