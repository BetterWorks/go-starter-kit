package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jasonsites/gosk-api/internal/resolver"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	run(ctx)
}

// run creates a new resolver with associated context group,
// then runs goroutines for serving http requests and graceful app shutdown
func run(c context.Context) {
	g, ctx := errgroup.WithContext(c)
	r := resolver.NewResolver(ctx, nil)

	// initialize the app resolver and starts the http server
	g.Go(func() error {
		if err := r.Initialize(); err != nil {
			return err
		}
		server, err := r.HTTPServer()
		if err != nil {
			return err
		}
		if err := server.Serve(); err != nil {
			return err
		}
		return nil
	})

	// gracefully shut down the http server and close the db connection pool
	g.Go(func() error {
		<-ctx.Done()

		fmt.Println("\nshutdown initiated")

		// shutdown server
		server, err := r.HTTPServer()
		if err != nil {
			return err
		}
		if err := server.App.Shutdown(); err != nil {
			return err
		}
		fmt.Println("server shut down")

		// close db pool
		pool, err := r.PostgreSQLClient()
		if err != nil {
			return err
		}
		pool.Close()
		fmt.Println("db connection pool closed")

		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
