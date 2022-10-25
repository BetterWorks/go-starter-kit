package resolver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jasonsites/gosk-api/config"
	"github.com/jasonsites/gosk-api/internal/httpapi"
	"github.com/sirupsen/logrus"
)

// Config defines the input to NewResolver
type Config struct {
	Config           *config.Configuration
	HTTPServer       *httpapi.Server
	Log              logrus.FieldLogger
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
	log              logrus.FieldLogger
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

	r.Metadata()
	r.Config()
	r.Log()

	return r
}

// Metadata provides a singleton application Metadata instance
func (r *Resolver) Metadata() *Metadata {
	if r.metadata == nil {
		var metadata *Metadata

		jsondata, err := os.ReadFile("/app/package.json")
		if err != nil {
			fmt.Printf("error reading package.json file, %v:", err)
		}

		if err := json.Unmarshal(jsondata, &metadata); err != nil {
			fmt.Printf("error unmarshalling package.json, %v:", err)
		}

		r.metadata = metadata
	}

	return r.metadata
}
