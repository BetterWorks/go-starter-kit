package resolver

import (
	"fmt"
	"log/slog"

	"github.com/BetterWorks/go-starter-kit/config"
	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
)

func logLevel(l string) slog.Leveler {
	switch l {
	case logger.LevelDebug:
		return slog.LevelDebug
	case logger.LevelInfo:
		return slog.LevelInfo
	case logger.LevelWarn:
		return slog.LevelWarn
	case logger.LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

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
