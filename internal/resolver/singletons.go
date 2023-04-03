package resolver

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"

	"github.com/BetterWorks/gosk-api/config"
	"github.com/BetterWorks/gosk-api/internal/domain"
	"github.com/BetterWorks/gosk-api/internal/httpapi"
	"github.com/BetterWorks/gosk-api/internal/repo"
	"github.com/BetterWorks/gosk-api/internal/types"
	"github.com/BetterWorks/gosk-api/internal/validation"
)

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

// Domain provides a singleton domain.Domain instance
func (r *Resolver) Domain() (*domain.Domain, error) {
	if r.domain == nil {
		svcResource, err := domain.NewResourceService(&domain.ResourceServiceConfig{
			Logger: &types.Logger{
				Enabled: r.config.Logger.SvcExample.Enabled,
				Level:   r.config.Logger.SvcExample.Level,
				Log:     r.log,
			},
			Repo: r.repoResource,
		})
		if err != nil {
			log.Printf("error resolving domain resource service: %v", err)
			return nil, err
		}

		services := &domain.Services{
			ResourceService: svcResource,
		}

		r.domain = domain.NewDomain(services)
	}

	return r.domain, nil
}

// HTTPServer provides a singleton httpapi.Server instance
func (r *Resolver) HTTPServer() (*httpapi.Server, error) {
	if r.httpServer == nil {
		server, err := httpapi.NewServer(&httpapi.Config{
			BaseURL: r.config.HttpAPI.BaseURL,
			Domain:  r.domain,
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

// RepositoryResource provides a singleton repo.resourceRepository instance
func (r *Resolver) RepositoryResource() (types.Repository, error) {
	if r.repoResource == nil {
		repo, err := repo.NewResourceRepository(&repo.ResourceRepoConfig{
			DBClient: r.postgreSQLClient,
			Logger: &types.Logger{
				Enabled: r.config.Logger.Repo.Enabled,
				Level:   r.config.Logger.Repo.Level,
				Log:     r.log,
			},
		})
		if err != nil {
			log.Printf("error resolving resource repository: %v", err)
			return nil, err
		}

		r.repoResource = repo
	}

	return r.repoResource, nil
}
