package resolver

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/rs/zerolog"

	"github.com/jasonsites/gosk-api/config"
	"github.com/jasonsites/gosk-api/internal/application"
	"github.com/jasonsites/gosk-api/internal/application/domain"
	"github.com/jasonsites/gosk-api/internal/application/services"
	"github.com/jasonsites/gosk-api/internal/httpapi"
	"github.com/jasonsites/gosk-api/internal/repo"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// Application provides a singleton application.Application instance
func (r *resolver) Application() *application.Application {
	if r.application == nil {
		services := &application.Services{
			EpisodeService: services.NewEpisodeService(r.repoEpisode),
			SeasonService:  services.NewSeasonService(r.repoSeason),
		}
		app := application.NewApplication(services)
		r.application = app
	}

	return r.application
}

// Config provides a singleton config.Configuration instance
func (r *resolver) Config() *config.Configuration {
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
func (r *resolver) HTTPServer() *httpapi.Server {
	if r.httpServer == nil {
		server, err := httpapi.NewServer(&httpapi.Config{
			Application: r.application,
			BaseURL:     r.config.HttpAPI.BaseURL,
			Logger: &domain.Logger{
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
func (r *resolver) Log() *zerolog.Logger {
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

// Metadata provides a singleton application Metadata instance
func (r *resolver) Metadata() *Metadata {
	if r.metadata == nil {
		var metadata *Metadata

		jsondata, err := os.ReadFile(r.config.Metadata.Path)
		if err != nil {
			log.Printf("error reading package.json file, %v:", err)
		}

		if err := json.Unmarshal(jsondata, &metadata); err != nil {
			log.Printf("error unmarshalling package.json, %v:", err)
		}

		r.metadata = metadata
	}

	return r.metadata
}

// PostgresClient provides a singleton postgres sql.DB instance
func (r *resolver) PostgreSQLClient() *sql.DB {
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

// RepositoryEpisode provides a singleton EpisodeRepository (interface) implementation
func (r *resolver) RepositoryEpisode() domain.EpisodeRepository {
	if r.repoEpisode == nil {

		repo, err := repo.NewEpisodeRepository(&repo.EpisodeRepoConfig{
			Logger: &domain.Logger{
				Enabled: r.config.Logger.Repo.Enabled,
				Level:   r.config.Logger.Repo.Level,
				Log:     r.log,
			},
			DBClient: r.postgreSQLClient,
		})
		if err != nil {
			log.Panicf("error resolving episode repository: %v", err)
		}

		r.repoEpisode = repo
	}

	return r.repoEpisode
}

// RepositorySeason provides a singleton SeasonRepository (interface) implementation
func (r *resolver) RepositorySeason() domain.SeasonRepository {
	if r.repoSeason == nil {

		repo, err := repo.NewSeasonRepository(&repo.SeasonRepoConfig{
			Logger: &domain.Logger{
				Enabled: r.config.Logger.Repo.Enabled,
				Level:   r.config.Logger.Repo.Level,
				Log:     r.log,
			},
			PostgreSQLClient: r.postgreSQLClient,
		})
		if err != nil {
			log.Panicf("error resolving season repository: %v", err)
		}

		r.repoSeason = repo
	}

	return r.repoSeason
}
