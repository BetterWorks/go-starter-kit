package main

import (
	"github.com/BetterWorks/go-starter-kit/internal/runtime"
)

func main() {
	runconf := &runtime.RunConfig{HTTPServer: true}
	runtime.NewRuntime(nil).Run(runconf)
}
