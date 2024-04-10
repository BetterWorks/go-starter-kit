package testutils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/BetterWorks/go-starter-kit/internal/http/httpserver"
	"github.com/BetterWorks/go-starter-kit/internal/resolver"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Cleanup deletes all rows on all database tables
func Cleanup(r *resolver.Resolver) error {
	db := r.PostgreSQLClient()

	tables := []string{"resource_entity"}

	for _, t := range tables {
		sql := fmt.Sprintf("DELETE from %s", t)
		_, err := db.Exec(context.TODO(), sql)
		if err != nil {
			return err
		}
	}

	return nil
}

// InitializeApp creates a new Resolver from the given config and returns a reference to the HTTP Server instance, the DB driver, and the Resolver itself
func InitializeApp(conf *resolver.Config) (*httpserver.Server, *pgxpool.Pool, *resolver.Resolver, error) {
	port := os.Getenv("HTTP_SERVER_PORT")
	if port == "" {
		os.Setenv("HTTP_SERVER_PORT", "9001")
	}
	r := resolver.NewResolver(context.Background(), conf)
	r.Load(resolver.LoadEntries.HTTPServer)
	server := r.HTTPServer()
	db := r.PostgreSQLClient()

	return server, db, r, nil
}

// SetRequestData creates a new HTTP Request instance from the given data
func SetRequestData(method, route string, body io.Reader, headers map[string]string) *http.Request {
	port := os.Getenv("HTTP_SERVER_PORT")
	if port == "" {
		os.Setenv("HTTP_SERVER_PORT", "9001")
		port = "9001"
	}
	url := fmt.Sprintf("http://localhost:%s%s", port, route)
	req, _ := http.NewRequest(method, url, body)
	if headers != nil {
		req = SetRequestHeaders(req, headers)
	}
	return req
}

// SetRequestHeaders set all headers on the given request
func SetRequestHeaders(req *http.Request, headers map[string]string) *http.Request {
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return req
}
