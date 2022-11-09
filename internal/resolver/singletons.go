package resolver

import (
	"database/sql"
	"log"
	"os"

	"github.com/rs/zerolog"

	"github.com/jasonsites/gosk-api/config"
	"github.com/jasonsites/gosk-api/internal/application"
	"github.com/jasonsites/gosk-api/internal/application/services"
	"github.com/jasonsites/gosk-api/internal/core/types"
	"github.com/jasonsites/gosk-api/internal/httpapi"
	"github.com/jasonsites/gosk-api/internal/repo/entities"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// Application provides a singleton application.Application instance
func (r *Resolver) Application() *application.Application {
	if r.application == nil {
		services := &application.Services{
			BookService:  services.NewBookService(entities.NewBookEntity()),
			MovieService: services.NewMovieService(entities.NewMovieEntity()),
		}
		app := application.NewApplication(services)
		r.application = app
	}

	return r.application
}

// Config provides a singleton config.Configuration instance
func (r *Resolver) Config() *config.Configuration {
	if r.config == nil {
		c, err := config.LoadConfiguration()
		if err != nil {
			log.Panicf("error resolving config: %v", err)
		}

		r.config = c
	}

	return r.config
}

// HTTPServer provides a singleton httpapi.Server instance
func (r *Resolver) HTTPServer() *httpapi.Server {
	if r.httpServer == nil {
		server, err := httpapi.NewServer(&httpapi.Config{
			Application: r.application,
			BaseURL:     r.config.HttpAPI.BaseURL,
			Logger: &types.Logger{
				Enabled: r.config.Logger.Http.Enabled,
				Level:   r.config.Logger.Http.Level,
				Log:     r.log,
			},
			Mode:      r.config.HttpAPI.Mode,
			Namespace: r.config.HttpAPI.Namespace,
			Port:      r.config.HttpAPI.Port,
		})
		if err != nil {
			log.Panicf("error resolving http server: %v", err)
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
			Str("name", r.metadata.Name).
			Str("version", r.metadata.Version).
			Timestamp().Logger()

		r.log = &logger
	}

	return r.log
}

// PostgresClient provides a singleton postgres sql.DB instance
func (r *Resolver) PostgreSQLClient() *sql.DB {
	if r.postgreSQLClient == nil {
		if err := validation.Validate.StructPartial(r.config, "Postgres"); err != nil {
			log.Panicf("invalid postgres config: %v", err)
		}

		db, err := sql.Open("postgres", postgresDSN(r.config.Postgres))
		if err != nil {
			log.Panicf("error resolving postgres client: %v", err)
		}

		r.postgreSQLClient = db
	}

	return r.postgreSQLClient
}

// BookRepository provides a singleton BookRepository (interface) implementation
func (r *Resolver) BookRepository() types.BookRepository {
	if r.bookRepository == nil {
		repo := entities.NewBookEntity()
		r.bookRepository = repo
	}

	return r.bookRepository
}

// MovieRepository provides a singleton MovieRepository (interface) implementation
func (r *Resolver) MovieRepository() types.MovieRepository {
	if r.movieRepository == nil {
		repo := entities.NewMovieEntity()
		r.movieRepository = repo
	}

	return r.movieRepository
}
