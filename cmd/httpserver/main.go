package main

import (
	"github.com/BetterWorks/go-starter-kit/config"
	"github.com/BetterWorks/go-starter-kit/internal/app"
	"github.com/BetterWorks/go-starter-kit/internal/runtime"
)

func main() {
	conf, err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	// Validate the configuration for HTTP server
	if err := app.Validator.Validate.Struct(conf.HTTP); err != nil {
		panic(err)
	}

	runconf := &runtime.RunConfig{HTTPServer: true}
	runtime.NewRuntime(nil).Run(runconf)
}
