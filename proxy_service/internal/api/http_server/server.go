package http_server

import (
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type HTTPServer struct {
	Mediator app.ProxyMediator
}

func (srv HTTPServer) subscribeToInbox(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
	}
	login := r.URL.Query().Get("login")
	pass := r.URL.Query().Get("password")
	srv.Mediator.Handle(contract.ProxySubCmd{Login: login, Pass: pass}, r.Context())
}
