package main

import (
	"github.com/BetterWorks/go-starter-kit/internal/runtime"
)

func main() {
	runconf := &runtime.RunConfig{Lambda: true}
	runtime.NewRuntime(nil).Run(runconf)
}
