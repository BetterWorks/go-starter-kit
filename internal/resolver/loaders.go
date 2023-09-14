package resolver

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BetterWorks/gosk-api/config"
	"github.com/BetterWorks/gosk-api/internal/core/interfaces"
	"github.com/BetterWorks/gosk-api/internal/core/logger"
	"github.com/BetterWorks/gosk-api/internal/core/query"
	"github.com/BetterWorks/gosk-api/internal/core/validation"
	"github.com/BetterWorks/gosk-api/internal/domain"
	"github.com/BetterWorks/gosk-api/internal/http/controllers"
	"github.com/BetterWorks/gosk-api/internal/http/httpserver"
	"github.com/BetterWorks/gosk-api/internal/repos"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Config provides a singleton config.Configuration instance
func (r *Resolver) Config() *config.Configuration {
	if r.config == nil {
		conf, err := config.LoadConfiguration()
		if err != nil {
			err = fmt.Errorf("config load error: %w", err)
			log.Error().Err(err).Send()
			panic(err)
		}

		r.config = conf
	}

	return r.config
}

// Domain provides a singleton domain.Domain instance
func (r *Resolver) Domain() *domain.Domain {
	if r.domain == nil {
		services := &domain.Services{
			Example: r.ExampleService(),
		}

		app, err := domain.NewDomain(services)
		if err != nil {
			err = fmt.Errorf("domain load error: %w", err)
			log.Error().Err(err).Send()
			panic(err)
		}

		r.domain = app
	}

	return r.domain
}

// ExampleRepository provides a singleton repo.exampleRepository instance
func (r *Resolver) ExampleRepository() interfaces.ExampleRepository {
	if r.exampleRepo == nil {
		repo, err := repos.NewExampleRepository(&repos.ExampleRepoConfig{
			DBClient: r.PostgreSQLClient(),
			Logger: &logger.Logger{
				Enabled: r.Config().Logger.Repo.Enabled,
				Level:   r.Config().Logger.Repo.Level,
				Log:     r.Log(),
			},
		})
		if err != nil {
			err = fmt.Errorf("example respository load error: %w", err)
			log.Error().Err(err).Send()
			panic(err)
		}

		r.exampleRepo = repo
	}

	return r.exampleRepo
}

// ExampleService provides a singleton domain.exampleService instance
func (r *Resolver) ExampleService() interfaces.Service {
	if r.exampleService == nil {
		svc, err := domain.NewExampleService(&domain.ExampleServiceConfig{
			Logger: &logger.Logger{
				Enabled: r.Config().Logger.Domain.Enabled,
				Level:   r.Config().Logger.Domain.Level,
				Log:     r.Log(),
			},
			Repo: r.ExampleRepository(),
		})
		if err != nil {
			err = fmt.Errorf("example service load error: %w", err)
			log.Error().Err(err).Send()
			panic(err)
		}

		r.exampleService = svc
	}

	return r.exampleService
}

// HTTPServer provides a singleton httpserver.Server instance
func (r *Resolver) HTTPServer() *httpserver.Server {
	if r.httpServer == nil {
		c := r.Config()

		queryConfig := func() *controllers.QueryConfig {
			limit := int(c.HTTP.API.Paging.DefaultLimit)
			offset := int(c.HTTP.API.Paging.DefaultOffset)

			attr := c.HTTP.API.Sorting.DefaultAttr
			order := c.HTTP.API.Sorting.DefaultOrder

			return &controllers.QueryConfig{
				Defaults: &controllers.QueryDefaults{
					Paging: &query.QueryPaging{
						Limit:  &limit,
						Offset: &offset,
					},
					Sorting: &query.QuerySorting{
						Attr:  &attr,
						Order: &order,
					},
				},
			}
		}()

		routeConfig := &httpserver.RouteConfig{Namespace: c.HTTP.Server.Namespace}

		server, err := httpserver.NewServer(&httpserver.ServerConfig{
			BaseURL: c.HTTP.Server.BaseURL,
			Domain:  r.Domain(),
			Logger: &logger.Logger{
				Enabled: c.Logger.HTTP.Enabled,
				Level:   c.Logger.HTTP.Level,
				Log:     r.Log(),
			},
			Mode:        c.HTTP.Server.Mode,
			Port:        c.HTTP.Server.Port,
			QueryConfig: queryConfig,
			RouteConfig: routeConfig,
		})
		if err != nil {
			err = fmt.Errorf("http server load error: %w", err)
			log.Error().Err(err).Send()
			panic(err)
		}

		r.httpServer = server
	}

	return r.httpServer
}

// Log provides a singleton zerolog.Logger instance
func (r *Resolver) Log() *zerolog.Logger {
	if r.log == nil {
		logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel).With().
			Int("pid", os.Getpid()).
			Str("name", r.Metadata().Name).
			Str("version", r.Metadata().Version).
			Timestamp().Logger()

		r.log = &logger
	}

	return r.log
}

// Metadata provides a singleton application Metadata instance
func (r *Resolver) Metadata() *Metadata {
	if r.metadata == nil {
		var metadata *Metadata

		jsondata, err := os.ReadFile(r.config.Metadata.Path)
		if err != nil {
			err = fmt.Errorf("package.json read error: %w", err)
			log.Error().Err(err).Send()
			panic(err)
		}

		if err := json.Unmarshal(jsondata, &metadata); err != nil {
			err = fmt.Errorf("package.json unmarshall error: %w", err)
			log.Error().Err(err).Send()
			panic(err)
		}

		r.metadata = metadata
	}

	return r.metadata
}

// PostgreSQLClient provides a singleton postgres pgxpool.Pool instance
func (r *Resolver) PostgreSQLClient() *pgxpool.Pool {
	if r.postgreSQLClient == nil {
		if err := validation.Validate.StructPartial(r.config, "Postgres"); err != nil {
			err = fmt.Errorf("invalid postgres config: %w", err)
			log.Error().Err(err).Send()
			panic(err)
		}

		client, err := pgxpool.New(r.appContext, postgresDSN(r.config.Postgres))
		if err != nil {
			err = fmt.Errorf("postgres client load error: %w", err)
			log.Error().Err(err).Send()
			panic(err)
		}

		r.postgreSQLClient = client
	}

	return r.postgreSQLClient
}
