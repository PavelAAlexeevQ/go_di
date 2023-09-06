//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"go_di/static/implementations"

	"github.com/google/wire"
)

func InitializeServer() *implementations.HTTPServer {
	wire.Build(implementations.ProvideHTTPServer, implementations.ProvideLogger)
	return &implementations.HTTPServer{}
}
