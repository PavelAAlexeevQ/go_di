package implementations

import (
	"fmt"
	"go_di/dynamic/interfaces"
	"io"
	"net/http"
	"sync"
	"time"

	"go.uber.org/dig"
)

type HTTPServer struct {
	logger                  interfaces.ILogger
	fastEndPointCalledCount int64
	slowEndPointCalledCount int64
	invokeDuration          time.Duration

	mu sync.Mutex

	diContainer *dig.Container
}

func (http *HTTPServer) FastEndPoint(w http.ResponseWriter, r *http.Request) {
	http.logger.Logf("called %v times\n", http.fastEndPointCalledCount)
	io.WriteString(w, "Invoke pre-constructed instance\n")
	http.fastEndPointCalledCount++
}

func (http *HTTPServer) SlowEndPoint(w http.ResponseWriter, r *http.Request) {
	http.mu.Lock()
	var callTime = time.Now()
	http.diContainer.Invoke(func(logger interfaces.ILogger) {
		logger.Logf("called %v times\n", http.slowEndPointCalledCount)
	})
	io.WriteString(w, "Get and invoke instance\n")
	var invocationTime = time.Now()
	var duration = invocationTime.Sub(callTime)
	http.slowEndPointCalledCount++
	http.invokeDuration += duration
	http.logger.Logf("Accumulated locked duration %v\n", http.invokeDuration.Seconds())
	http.mu.Unlock()
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
		logger:                  log,
		slowEndPointCalledCount: 1,
		fastEndPointCalledCount: 1,
		diContainer:             container,
	}
}
