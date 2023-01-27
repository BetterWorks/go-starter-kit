package main

import (
	"log"

	"github.com/jasonsites/gosk-api/internal/resolver"
)

func main() {
	defer recovery()
	r := resolver.NewResolver(nil)
	r.HTTPServer().Serve()
}

func recovery() {
	if err := recover(); err != nil {
		log.Fatalf("app recovery failed: %v", err)
	}
}
