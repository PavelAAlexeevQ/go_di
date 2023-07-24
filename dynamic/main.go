package main

import (
	"go.uber.org/dig"

	"go_di/dynamic/implementations"
	"go_di/dynamic/interfaces"
)

func SetupDI() (container *dig.Container) {
	container = dig.New()

	var err error = nil

	err = container.Provide(func() interfaces.ILogger {
		return implementations.ProvideLogger()
	})
	if err != nil {
		panic(err)
	}

	err = container.Provide(func(logger interfaces.ILogger) interfaces.IHTTPServer {
		return implementations.ProvideHTTPServer(logger, container)
	})
	if err != nil {
		panic(err)
	}
	return container
}

func main() {
	container := SetupDI()

	err := container.Invoke(func(httpServer interfaces.IHTTPServer) {
		httpServer.SetupHTTPServer()
	})
	if err != nil {
		panic(err)
	}
}
