package resolver

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/jasonsites/gosk-api/config"
	"github.com/jasonsites/gosk-api/internal/httpapi"
	"github.com/rs/zerolog"
)

// Config defines the input to NewResolver
type Config struct {
	Config           *config.Configuration
	HTTPServer       *httpapi.Server
	Log              *zerolog.Logger
	Metadata         *Metadata
	PostgreSQLClient *sql.DB
}

// Application metadata
type Metadata struct {
	Name    string
	Version string
}

// Resolver provides singleton instances of app components
type Resolver struct {
	config           *config.Configuration
	httpServer       *httpapi.Server
	log              *zerolog.Logger
	metadata         *Metadata
	postgreSQLClient *sql.DB
}

// NewResolver returns a new Resolver instance
func NewResolver(c *Config) *Resolver {
	if c == nil {
		c = &Config{}
	}

	r := &Resolver{
		config:           c.Config,
		httpServer:       c.HTTPServer,
		log:              c.Log,
		metadata:         c.Metadata,
		postgreSQLClient: c.PostgreSQLClient,
	}

	r.initialize()

	return r
}

func (r *Resolver) initialize() {
	r.Config()
	r.Metadata()
	r.Log()
	r.PostgreSQLClient()
	r.HTTPServer()
}

// Metadata provides a singleton application Metadata instance
func (r *Resolver) Metadata() *Metadata {
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