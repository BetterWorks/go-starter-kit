package resolver

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/BetterWorks/go-starter-kit/config"
	"github.com/BetterWorks/go-starter-kit/internal/core/app"
	"github.com/BetterWorks/go-starter-kit/internal/core/interfaces"
	"github.com/BetterWorks/go-starter-kit/internal/domain"
	"github.com/BetterWorks/go-starter-kit/internal/http/httpserver"
	"github.com/BetterWorks/go-starter-kit/internal/lambda"
	"github.com/jackc/pgx/v5/pgxpool"
	ld "github.com/launchdarkly/go-server-sdk/v7"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Application metadata
type Metadata struct {
	Name    string
	Version string
}

// Config defines the input to NewResolver
type Config struct {
	Config           *config.Configuration
	Domain           *domain.Domain
	ExampleRepo      interfaces.ExampleRepository
	ExampleService   interfaces.ExampleService
	FlagsClient      *ld.LDClient
	HTTPServer       *httpserver.Server
	LambdaService    *lambda.LambdaService
	Log              *slog.Logger
	Metadata         *app.Metadata
	NewRelicClient   *newrelic.Application
	PostgreSQLClient *pgxpool.Pool
}

// Resolver provides a configurable app component graph
type Resolver struct {
	appContext       context.Context
	config           *config.Configuration
	domain           *domain.Domain
	exampleRepo      interfaces.ExampleRepository
	exampleService   interfaces.ExampleService
	flagsClient      *ld.LDClient
	httpServer       *httpserver.Server
	lambdaService    *lambda.LambdaService
	log              *slog.Logger
	metadata         *app.Metadata
	newRelicClient   *newrelic.Application
	postgreSQLClient *pgxpool.Pool
}

// NewResolver returns a new Resolver instance
func NewResolver(ctx context.Context, c *Config) *Resolver {
	if c == nil {
		c = &Config{}
	}

	r := &Resolver{
		appContext:       ctx,
		config:           c.Config,
		domain:           c.Domain,
		exampleRepo:      c.ExampleRepo,
		exampleService:   c.ExampleService,
		flagsClient:      c.FlagsClient,
		httpServer:       c.HTTPServer,
		log:              c.Log,
		lambdaService:    c.LambdaService,
		metadata:         c.Metadata,
		newRelicClient:   c.NewRelicClient,
		postgreSQLClient: c.PostgreSQLClient,
	}

	return r
}

// LoadEntries provides option strings for loading the resolver from various entry nodes
// in the app component graph (cli, http, lambda)
var LoadEntries = struct {
	HTTPServer string
	Lambda     string
}{HTTPServer: "http", Lambda: "lambda"}

// Load resolves app components starting from the given entry node of the component graph
func (r *Resolver) Load(entry string) {
	switch entry {
	case LoadEntries.HTTPServer:
		r.HTTPServer()
	case LoadEntries.Lambda:
		r.LambdaService()
	default:
		panic(fmt.Errorf("invalid resolver load entry point '%s'", entry))
	}
}
