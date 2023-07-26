package implementations

import (
	"fmt"
	"go_di/static/interfaces"
	"io"
	"net/http"
)

type HTTPServer struct {
	logger              interfaces.ILogger
	endPointCalledCount int64
}

// race condition!
func (http *HTTPServer) EndPoint(w http.ResponseWriter, r *http.Request) {
	http.logger.Logf("called %v times\n", http.endPointCalledCount)
	io.WriteString(w, "Invoke pre-constructed instance\n")
	http.endPointCalledCount++
}

func (httpServer *HTTPServer) SetupHTTPServer() {
	http.HandleFunc("/endPoint", httpServer.EndPoint)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		httpServer.logger.Log(err.Error())
	}
}

func ProvideHTTPServer(log interfaces.ILogger) *HTTPServer {
	fmt.Println("ProvideHTTPServer()")
	return &HTTPServer{
		logger:              log,
		endPointCalledCount: 1,
	}
}
