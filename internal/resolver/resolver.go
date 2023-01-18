package resolver

import (
	"database/sql"

	"github.com/jasonsites/gosk-api/config"
	"github.com/jasonsites/gosk-api/internal/application"
	"github.com/jasonsites/gosk-api/internal/application/domain"
	"github.com/jasonsites/gosk-api/internal/httpapi"
	"github.com/rs/zerolog"
)

// Config defines the input to NewResolver
type Config struct {
	Application      *application.Application
	Config           *config.Configuration
	HTTPServer       *httpapi.Server
	Log              *zerolog.Logger
	Metadata         *Metadata
	PostgreSQLClient *sql.DB
	RepoEpisode      domain.Repository
	RepoSeason       domain.Repository
}

// Application metadata
type Metadata struct {
	Name    string
	Version string
}

// Resolver provides singleton instances of app components
type resolver struct {
	application      *application.Application
	config           *config.Configuration
	httpServer       *httpapi.Server
	log              *zerolog.Logger
	metadata         *Metadata
	postgreSQLClient *sql.DB
	repoEpisode      domain.Repository
	repoSeason       domain.Repository
}

// NewResolver returns a new Resolver instance
func NewResolver(c *Config) *resolver {
	if c == nil {
		c = &Config{}
	}

	r := &resolver{
		application:      c.Application,
		config:           c.Config,
		httpServer:       c.HTTPServer,
		log:              c.Log,
		metadata:         c.Metadata,
		postgreSQLClient: c.PostgreSQLClient,
		repoEpisode:      c.RepoEpisode,
		repoSeason:       c.RepoSeason,
	}

	r.initialize()

	return r
}

// initialize
func (r *resolver) initialize() {
	r.Config()
	r.Metadata()
	r.Log()
	r.PostgreSQLClient()
	r.RepositoryEpisode()
	r.RepositorySeason()
	r.Application()
	r.HTTPServer()
}
