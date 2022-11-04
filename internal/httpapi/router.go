package httpapi

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	mw "github.com/jasonsites/gosk-api/internal/httpapi/middleware"
	"github.com/jasonsites/gosk-api/internal/httpapi/routes"
)

// configureMiddleware
func (s *Server) configureMiddleware() {
	r := s.Router

	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(mw.ResponseLogger(s.Logger))
	// security headers
	// "github.com/gin-contrib/cors"
	r.Use(mw.ErrorHandler())
	r.Use(mw.Correlation())
	r.Use(mw.RequestLogger(s.Logger))
}

// registerRoutes
func (s *Server) registerRoutes() {
	r := s.Router
	c := s.controller
	ns := s.namespace

	routes.BaseRouter(r, c, ns)
	routes.HealthRouter(r, c, ns)

	routes.BookRouter(r, c, ns)
	routes.MovieRouter(r, c, ns)
}
