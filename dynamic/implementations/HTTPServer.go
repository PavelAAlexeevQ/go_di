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

func (http *HTTPServer) HelloEndPoint(w http.ResponseWriter, r *http.Request) {
	http.logger.Logf("called %v times\n", http.helloCalledCount)
	io.WriteString(w, "Hello from DI\n")
	http.helloCalledCount++
}

func (http *HTTPServer) SpeedTest(w http.ResponseWriter, r *http.Request) {
	http.diContainer.Invoke(func(logger interfaces.ILogger) {
		logger.Logf("called %v times\n", http.speedTestCalledCount)
	})
	io.WriteString(w, "Hello from DI\n")
	http.speedTestCalledCount++
}

func (httpServer *HTTPServer) SetupHTTPServer() {
	http.HandleFunc("/hello", httpServer.HelloEndPoint)

	http.HandleFunc("/speedTest", httpServer.SpeedTest)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		httpServer.logger.Log(err.Error())
	}
}

func ProvideHTTPServer(log interfaces.ILogger, container *dig.Container) interfaces.IHTTPServer {
	fmt.Println("ProvideHTTPServer()")
	return &HTTPServer{
		logger:               log,
		speedTestCalledCount: 0,
		diContainer:          container,
	}
}
