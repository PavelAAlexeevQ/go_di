package main

import (
	"go_di/static/interfaces"
)

type ServerImpl struct {
	interfaces.IHTTPServer
}

func main() {
	server := InitializeServer()
	server.SetupHTTPServer()
}
