package implementations

import (
	"fmt"
	"go_di/dynamic/interfaces"
	"io"
	"net/http"
	"sync"

	"go.uber.org/dig"
)

type HTTPServer struct {
	logger               interfaces.ILogger
	endPoint1CalledCount int64
	endPoint2CalledCount int64

	mu1 sync.Mutex
	mu2 sync.Mutex

	diContainer *dig.Container
}

func (http *HTTPServer) EndPoint1(w http.ResponseWriter, r *http.Request) {
	http.mu1.Lock()
	defer http.mu1.Unlock()

	http.logger.Logf("called %v times\n", http.endPoint1CalledCount)
	io.WriteString(w, "Invoke pre-constructed instance\n")
	http.endPoint1CalledCount++
}

func (http *HTTPServer) EndPoint2(w http.ResponseWriter, r *http.Request) {
	http.mu2.Lock()
	defer http.mu2.Unlock()

	http.diContainer.Invoke(func(logger interfaces.ILogger) {
		logger.Logf("called %v times\n", http.endPoint2CalledCount)
	})
	io.WriteString(w, "Get and invoke instance\n")
	http.endPoint2CalledCount++
}

func (httpServer *HTTPServer) SetupHTTPServer() {
	http.HandleFunc("/EndPoint1", httpServer.EndPoint1)

	http.HandleFunc("/EndPoint2", httpServer.EndPoint2)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		httpServer.logger.Log(err.Error())
	}
}

func ProvideHTTPServer(log interfaces.ILogger, container *dig.Container) interfaces.IHTTPServer {
	fmt.Println("ProvideHTTPServer()")
	return &HTTPServer{
		logger:               log,
		endPoint1CalledCount: 1,
		endPoint2CalledCount: 1,
		diContainer:          container,
	}
}
