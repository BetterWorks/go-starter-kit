package resolver

import (
	"database/sql"
	"log"
	"os"

	"github.com/rs/zerolog"

	"github.com/jasonsites/gosk-api/config"
	"github.com/jasonsites/gosk-api/internal/core"
	"github.com/jasonsites/gosk-api/internal/httpapi"
	"github.com/jasonsites/gosk-api/internal/validation"
)

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
			BaseURL: r.config.HttpAPI.BaseURL,
			Logger: &core.Logger{
				Enabled: r.config.Logger.Http.Enabled,
				Level:   r.config.Logger.Http.Level,
				Log:     r.log,
			},
			Namespace: r.config.HttpAPI.Namespace,
			Port:      r.config.HttpAPI.Port,
		})
		if err != nil {
			log.Panicf("error resolving grpc server: %v", err)
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
