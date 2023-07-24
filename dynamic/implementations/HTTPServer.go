package implementations

import (
	"fmt"
	"go_di/dynamic/interfaces"
	"io"
	"net/http"

	"go.uber.org/dig"
)

type HTTPServer struct {
	logger               interfaces.ILogger
	helloCalledCount     int64
	speedTestCalledCount int64
	diContainer          *dig.Container
}

func (http *HTTPServer) FastEndPoint(w http.ResponseWriter, r *http.Request) {
	http.logger.Logf("called %v times\n", http.helloCalledCount)
	io.WriteString(w, "Invoke pre-constructed instance\n")
	http.helloCalledCount++
}

func (http *HTTPServer) SlowEndPoint(w http.ResponseWriter, r *http.Request) {
	http.diContainer.Invoke(func(logger interfaces.ILogger) {
		logger.Logf("called %v times\n", http.speedTestCalledCount)
	})
	io.WriteString(w, "Get and invoke instance\n")
	http.speedTestCalledCount++
}

func (httpServer *HTTPServer) SetupHTTPServer() {
	http.HandleFunc("/fastEndPoint", httpServer.FastEndPoint)

	http.HandleFunc("/slowEndPoint", httpServer.SlowEndPoint)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		httpServer.logger.Log(err.Error())
	}
}

func ProvideHTTPServer(log interfaces.ILogger, container *dig.Container) interfaces.IHTTPServer {
	fmt.Println("ProvideHTTPServer()")
	return &HTTPServer{
		logger:               log,
		speedTestCalledCount: 1,
		helloCalledCount:     1,
		diContainer:          container,
	}
}
