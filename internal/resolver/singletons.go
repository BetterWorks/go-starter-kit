package resolver

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"github.com/jasonsites/gosk-api/config"
	"github.com/jasonsites/gosk-api/internal/application"
	"github.com/jasonsites/gosk-api/internal/httpapi"
	"github.com/jasonsites/gosk-api/internal/repo"
	"github.com/jasonsites/gosk-api/internal/types"
	"github.com/jasonsites/gosk-api/internal/validation"
)

// Application provides a singleton application.Application instance
func (r *Resolver) Application() (*application.Application, error) {
	if r.application == nil {
		svcEpisode, err := application.NewEpisodeService(&application.EpisodeServiceConfig{
			Logger: &types.Logger{
				Enabled: r.config.Logger.SvcExample.Enabled,
				Level:   r.config.Logger.SvcExample.Level,
				Log:     r.log,
			},
			Repo: r.repoEpisode,
		})
		if err != nil {
			log.Printf("error resolving application episode service: %v", err)
			return nil, err
		}

		svcSeason, err := application.NewSeasonService(&application.SeasonServiceConfig{
			Logger: &types.Logger{
				Enabled: r.config.Logger.SvcExample.Enabled,
				Level:   r.config.Logger.SvcExample.Level,
				Log:     r.log,
			},
			Repo: r.repoSeason,
		})
		if err != nil {
			log.Printf("error resolving application season service: %v", err)
			return nil, err
		}

		services := &application.Services{
			EpisodeService: svcEpisode,
			SeasonService:  svcSeason,
		}

		r.application = application.NewApplication(services)
	}

	return r.application, nil
}

// Config provides a singleton config.Configuration instance
func (r *Resolver) Config() (*config.Configuration, error) {
	if r.config == nil {
		c, err := config.LoadConfiguration()
		if err != nil {
			log.Printf("error resolving config: %v", err)
			return nil, err
		}
		r.config = c
	}

	return r.config, nil
}

// HTTPServer provides a singleton httpapi.Server instance
func (r *Resolver) HTTPServer() (*httpapi.Server, error) {
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
			log.Printf("error resolving http server: %v", err)
			return nil, err
		}
		r.httpServer = server
	}

	return r.httpServer, nil
}

// Log provides a singleton zerolog.Logger instance
func (r *Resolver) Log() (*zerolog.Logger, error) {
	if r.log == nil {
		logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel).With().
			Int("pid", os.Getpid()).
			Str("name", r.metadata.Name).
			Str("version", r.metadata.Version).
			Timestamp().Logger()

		r.log = &logger
	}

	return r.log, nil
}

// Metadata provides a singleton application Metadata instance
func (r *Resolver) Metadata() (*Metadata, error) {
	if r.metadata == nil {
		var metadata *Metadata

		jsondata, err := os.ReadFile(r.config.Metadata.Path)
		if err != nil {
			log.Printf("error reading package.json file, %v:", err)
			return nil, err
		}

		if err := json.Unmarshal(jsondata, &metadata); err != nil {
			log.Printf("error unmarshalling package.json, %v:", err)
			return nil, err
		}

		r.metadata = metadata
	}

	return r.metadata, nil
}

// PostgreSQLClient provides a singleton postgres pgxpool.Pool instance
func (r *Resolver) PostgreSQLClient() (*pgxpool.Pool, error) {
	if r.postgreSQLClient == nil {
		if err := validation.Validate.StructPartial(r.config, "Postgres"); err != nil {
			log.Printf("invalid postgres config: %v", err)
			return nil, err
		}

		client, err := pgxpool.New(r.context, postgresDSN(r.config.Postgres))
		if err != nil {
			log.Printf("error resolving postgres client: %v", err)
			return nil, err
		}

		r.postgreSQLClient = client
	}

	return r.postgreSQLClient, nil
}

// RepositoryEpisode provides a singleton repo.episodeRepository instance
func (r *Resolver) RepositoryEpisode() (types.Repository, error) {
	if r.repoEpisode == nil {
		repo, err := repo.NewEpisodeRepository(&repo.EpisodeRepoConfig{
			DBClient: r.postgreSQLClient,
			Logger: &types.Logger{
				Enabled: r.config.Logger.Repo.Enabled,
				Level:   r.config.Logger.Repo.Level,
				Log:     r.log,
			},
		})
		if err != nil {
			log.Printf("error resolving episode repository: %v", err)
			return nil, err
		}

		r.repoEpisode = repo
	}

	return r.repoEpisode, nil
}

// RepositorySeason provides a singleton repo.seasonRepository instance
func (r *Resolver) RepositorySeason() (types.Repository, error) {
	if r.repoSeason == nil {
		repo, err := repo.NewSeasonRepository(&repo.SeasonRepoConfig{
			DBClient: r.postgreSQLClient,
			Logger: &types.Logger{
				Enabled: r.config.Logger.Repo.Enabled,
				Level:   r.config.Logger.Repo.Level,
				Log:     r.log,
			},
		})
		if err != nil {
			log.Printf("error resolving season repository: %v", err)
			return nil, err
		}

		r.repoSeason = repo
	}

	return r.repoSeason, nil
}
