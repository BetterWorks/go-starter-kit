package main

import (
	"fmt"

	"github.com/jasonsites/gosk-api/internal/resolver"
	_ "github.com/lib/pq"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("app recovery failed: %v", err))
		}
	}()
	r := resolver.NewResolver(nil)
	r.HTTPServer().Serve()
}
