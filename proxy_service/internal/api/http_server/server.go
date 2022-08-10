package http_server

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app"
)

type httpServer struct {
	mediator app.Mediator
}

func newHTTPServer() *httpServer {
	return &httpServer{mediator: app.NewProxyMediator()}
}
