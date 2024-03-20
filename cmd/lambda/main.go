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

	// Validate the configuration for Lambda
	if err := app.Validator.Validate.Struct(conf.Lambda); err != nil {
		panic(err)
	}

	runconf := &runtime.RunConfig{Lambda: true}
	runtime.NewRuntime(nil).Run(runconf)
}
