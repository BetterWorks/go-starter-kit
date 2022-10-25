package resolver

import (
	"database/sql"
	"fmt"

	"github.com/jasonsites/gosk-api/config"
	"github.com/jasonsites/gosk-api/internal/httpapi"
	"github.com/jasonsites/gosk-api/internal/validation"
	"github.com/sirupsen/logrus"
)

// Config provides a singleton config.Configuration instance
func (r *Resolver) Config() *config.Configuration {
	if r.config == nil {
		c, err := config.LoadConfiguration()
		if err != nil {
			panic(fmt.Errorf("error resolving config: %v", err))
		}

		r.config = c
	}

	return r.config
}

// HTTPServer provides a singleton httpapi.Server instance
func (r *Resolver) HTTPServer() *httpapi.Server {
	if r.httpServer == nil {
		server, err := httpapi.NewServer(&httpapi.Config{
			BaseURL:   r.Config().HttpAPI.BaseURL,
			Log:       r.Log(),
			Namespace: r.Config().HttpAPI.Namespace,
			Port:      r.Config().HttpAPI.Port,
		})
		if err != nil {
			panic(fmt.Errorf("error resolving grpc server: %v", err))
		}

		r.httpServer = server
	}

	return r.httpServer
}

// Log provides a singleton logrus.FieldLogger instance
func (r *Resolver) Log() logrus.FieldLogger {
	if r.log == nil {
		r.log = logrus.WithFields(logrus.Fields{
			"name":    r.metadata.Name,
			"version": r.metadata.Version,
		})

		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.999Z07:00",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyLevel: "plevel",
			},
		})

		level, err := logrus.ParseLevel(r.Config().Logger.App.Level)
		if err != nil {
			level = logrus.InfoLevel
		}
		logrus.SetLevel(level)
	}

	return r.log
}

// PostgresClient provides a singleton postgres sql.DB instance
func (r *Resolver) PostgreSQLClient() *sql.DB {
	if r.postgreSQLClient == nil {
		if err := validation.Validate.StructPartial(r.Config(), "Postgres"); err != nil {
			panic(fmt.Errorf("invalid postgres config: %v", err))
		}

		db, err := sql.Open("postgres", postgresDSN(r.config.Postgres))
		if err != nil {
			panic(fmt.Errorf("error resolving postgres client: %v", err))
		}

		r.postgreSQLClient = db
	}

	return r.postgreSQLClient
}
