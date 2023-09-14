package main

import (
	"github.com/BetterWorks/gosk-api/internal/runtime"
)

func main() {
	runconf := &runtime.RunConfig{HTTPServer: true}
	runtime.NewRuntime(nil).Run(runconf)
}
