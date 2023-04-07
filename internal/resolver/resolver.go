package resolver

import (
	"context"

	"github.com/BetterWorks/gosk-api/config"
	"github.com/BetterWorks/gosk-api/internal/domain"
	"github.com/BetterWorks/gosk-api/internal/httpapi"
	"github.com/BetterWorks/gosk-api/internal/types"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

// Config defines the input to NewResolver
type Config struct {
	Config           *config.Configuration
	Domain           *domain.Domain
	HTTPServer       *httpapi.Server
	Log              *zerolog.Logger
	Metadata         *Metadata
	PostgreSQLClient *pgxpool.Pool
	RepoResource     types.Repository
}

// Application metadata
type Metadata struct {
	Name    string
	Version string
}

// Resolver provides singleton instances of app components
type Resolver struct {
	config           *config.Configuration
	context          context.Context
	domain           *domain.Domain
	httpServer       *httpapi.Server
	log              *zerolog.Logger
	metadata         *Metadata
	postgreSQLClient *pgxpool.Pool
	repoResource     types.Repository
}

// NewResolver returns a new Resolver instance
func NewResolver(ctx context.Context, c *Config) *Resolver {
	if c == nil {
		c = &Config{}
	}

	r := &Resolver{
		config:           c.Config,
		context:          ctx,
		domain:           c.Domain,
		httpServer:       c.HTTPServer,
		log:              c.Log,
		metadata:         c.Metadata,
		postgreSQLClient: c.PostgreSQLClient,
		repoResource:     c.RepoResource,
	}

	return r
}

// initialize bootstraps the application in dependency order
func (r *Resolver) Initialize() error {
	if _, err := r.Config(); err != nil {
		return err
	}
	if _, err := r.Metadata(); err != nil {
		return err
	}
	if _, err := r.Log(); err != nil {
		return err
	}
	if _, err := r.PostgreSQLClient(); err != nil {
		return err
	}
	if _, err := r.RepositoryResource(); err != nil {
		return err
	}
	if _, err := r.Domain(); err != nil {
		return err
	}
	if _, err := r.HTTPServer(); err != nil {
		return err
	}

	return nil
}
