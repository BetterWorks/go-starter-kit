package runtime

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/BetterWorks/gosk-api/internal/resolver"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

// Runtime provides an abstraction over the Resolver for running the application in various modes
// and handling graceful shutdown via goroutines
type Runtime struct {
	config *resolver.Config
}

// NewRuntime provides a new Runtime instance
func NewRuntime(c *resolver.Config) *Runtime {
	return &Runtime{config: c}
}

// RunConfig provides configuration options for running the application in various modes
// WARNING: only one option should be enabled per build/process
type RunConfig struct {
	HTTPServer bool
}

// Run creates a new Resolver with associated context group, then runs goroutines for bootstrapping
// the application and handling graceful shutdown
func (rt *Runtime) Run(conf *RunConfig) *resolver.Resolver {
	c, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	g, ctx := errgroup.WithContext(c)

	log.Info().Msg("initializing resolver")
	r := resolver.NewResolver(ctx, rt.config)

	// load resolver app components and start the configured application
	g.Go(func() error {
		if conf.HTTPServer {
			log.Info().Msg("loading resolver app components")
			r.Load(resolver.LoadEntries.HTTPServer)

			log.Info().Msg("starting http server")
			server := r.HTTPServer()
			if err := server.Serve(); err != nil {
				return err
			}
		}

		return nil
	})

	// gracefully shut down the configured application and close the db connection pool
	g.Go(func() error {
		<-ctx.Done()

		log.Info().Msg("shutdown initiated")

		if conf.HTTPServer {
			server := r.HTTPServer()
			if err := server.Server.Shutdown(context.Background()); err != nil {
				return err
			}
			log.Info().Msg("http server shut down")
		}

		pool := r.PostgreSQLClient()
		pool.Close()
		log.Info().Msg("db connection pool closed")

		log.Info().Msg("shutdown complete")

		return nil
	})

	if err := g.Wait(); err != nil {
		err = fmt.Errorf("application run error: %w", err)
		log.Error().Err(err).Send()
	}

	return r
}
