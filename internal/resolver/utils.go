package resolver

import (
	"fmt"

	"github.com/BetterWorks/gosk-api/config"
)

// postgresDSN returns a data source name string from a given postgres configuration
func postgresDSN(c config.Postgres) string {
	// TODO: sslmode
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Database,
	)
}
