package http_server

import (
	"github.com/Guise322/ozon-exercises/common/mediator"
)

type httpServer struct {
	mediator mediator.Mediator
}

func newHTTPServer(med mediator.Mediator) *httpServer {
	return &httpServer{mediator: med}
}
