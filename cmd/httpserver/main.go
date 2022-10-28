package main

import (
	"log"

	"github.com/jasonsites/gosk-api/internal/resolver"
	_ "github.com/lib/pq"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("app recovery failed: %v", err)
		}
	}()
	r := resolver.NewResolver(nil)
	r.HTTPServer().Serve()
}
