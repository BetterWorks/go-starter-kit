package httpserver

import (
	"compress/gzip"
	"fmt"
	"net/http"

	"github.com/BetterWorks/gosk-api/internal/core/logger"
	"github.com/BetterWorks/gosk-api/internal/domain"
	ctrl "github.com/BetterWorks/gosk-api/internal/http/controllers"
	mw "github.com/BetterWorks/gosk-api/internal/http/middleware"
	"github.com/BetterWorks/gosk-api/internal/http/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/goddtriffin/helmet"
)

type RouteConfig struct {
	Namespace string `validate:"required"`
}

type controllerRegistry struct {
	ExampleController *ctrl.Controller
}

// configureMiddleware
func configureMiddleware(r *chi.Mux, conf *RouteConfig, logger *logger.Logger) {
	skipHealth := func(r *http.Request) bool {
		return r.URL.Path == fmt.Sprintf("/%s/health", conf.Namespace)
	}

	r.Use(middleware.Compress(gzip.DefaultCompression))
	r.Use(mw.Correlation(&mw.CorrelationConfig{Next: skipHealth}))
	r.Use(mw.ResponseLogger(&mw.ResponseLoggerConfig{Logger: logger, Next: skipHealth}))
	r.Use(helmet.Default().Secure)
	// r.Use(middleware.RealIP)
	r.Use(mw.RequestLogger(&mw.RequestLoggerConfig{Logger: logger, Next: skipHealth}))
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

// registerControllers
func registerControllers(services *domain.Services, logger *logger.Logger, qc *ctrl.QueryConfig) *controllerRegistry {
	return &controllerRegistry{
		ExampleController: ctrl.NewController(&ctrl.Config{
			QueryConfig: qc,
			Logger:      logger,
			Service:     services.Example,
		}),
	}
}

// registerRoutes
func registerRoutes(r *chi.Mux, c *controllerRegistry, conf *RouteConfig) {
	ns := conf.Namespace
	routes.BaseRouter(r, nil, ns)
	routes.HealthRouter(r, nil, ns)
	routes.ExampleRouter(r, c.ExampleController, ns)
}
