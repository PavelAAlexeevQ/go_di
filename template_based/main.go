package main

import (
	"go_di/do/implementations"
	"go_di/do/interfaces"

	"github.com/samber/do"
)

func main() {
	injector := do.New()

	do.Provide(injector, implementations.ProvideLogger)
	do.Provide(injector, implementations.ProvideHTTPServer)

	server := do.MustInvoke[interfaces.IHTTPServer](injector)
	server.SetupHTTPServer()
}
