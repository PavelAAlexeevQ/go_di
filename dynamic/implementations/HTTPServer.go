package implementations

import (
	"fmt"
	"go_di/dynamic/interfaces"
	"io"
	"net/http"
)

type HTTPServer struct {
	logger      interfaces.ILogger
	calledCount int64
}

func (http *HTTPServer) HelloEndPoint(w http.ResponseWriter, r *http.Request) {
	http.logger.Logf("called %v times\n", http.calledCount)
	io.WriteString(w, "Hello from DI\n")
	http.calledCount++
}

func (httpServer *HTTPServer) SetupHTTPServer() {
	http.HandleFunc("/hello", httpServer.HelloEndPoint)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		httpServer.logger.Log(err.Error())
	}
}

func ProvideHTTPServer(log interfaces.ILogger) interfaces.IHTTPServer {
	fmt.Println("ProvideLogger()")
	return &HTTPServer{
		logger:      log,
		calledCount: 0,
	}
}
