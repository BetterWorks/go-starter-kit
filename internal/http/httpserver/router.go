package httpserver

import (
	"compress/gzip"
	"fmt"
	"net/http"

	"github.com/BetterWorks/go-starter-kit/internal/core/interfaces"
	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	"github.com/BetterWorks/go-starter-kit/internal/domain"
	ctrl "github.com/BetterWorks/go-starter-kit/internal/http/controllers"
	mw "github.com/BetterWorks/go-starter-kit/internal/http/middleware"
	"github.com/BetterWorks/go-starter-kit/internal/http/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/goddtriffin/helmet"
)

type RouterConfig struct {
	Namespace string `validate:"required"`
}

type controllerRegistry struct {
	ExampleController interfaces.ExampleController
}

// configureMiddleware
func configureMiddleware(conf *RouterConfig, r *chi.Mux, logger *logger.CustomLogger) {
	skipHealth := func(r *http.Request) bool {
		return r.URL.Path == fmt.Sprintf("/%s/health", conf.Namespace)
	}

	r.Use(middleware.Compress(gzip.DefaultCompression))
	r.Use(mw.Correlation(&mw.CorrelationConfig{Next: skipHealth}))
	r.Use(mw.ResponseLogger(&mw.ResponseLoggerConfig{Logger: logger, Next: skipHealth}))
	r.Use(helmet.Default().Secure)
	// r.Use(middleware.RealIP)
	r.Use(mw.RequestLogger(&mw.RequestLoggerConfig{Logger: logger, Next: skipHealth}))
	r.Use(mw.NotFound)
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
func registerControllers(services *domain.Services, logger *logger.CustomLogger, qc *ctrl.QueryConfig) (*controllerRegistry, error) {
	controller, err := ctrl.NewExampleController(&ctrl.ExampleControllerConfig{
		QueryConfig: qc,
		Logger:      logger,
		Service:     services.Example,
	})
	if err != nil {
		return nil, err
	}

	registry := &controllerRegistry{
		ExampleController: controller,
	}

	return registry, nil
}

// registerRoutes
func registerRoutes(conf *RouterConfig, r *chi.Mux, c *controllerRegistry) {
	ns := conf.Namespace
	routes.BaseRouter(r, ns)
	routes.HealthRouter(r, ns)
	routes.ExampleRouter(r, ns, c.ExampleController)
}
