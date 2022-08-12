package http_server

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
)

type httpServer struct {
	mediator app.Mediator
}

func newHTTPServer(cl interf.SubClient) *httpServer {
	return &httpServer{mediator: app.NewProxyMediator(cl)}
}
