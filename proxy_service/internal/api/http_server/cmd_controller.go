package http_server

import (
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type cmdController struct {
	mediator app.ProxyMediator
}

func (c cmdController) subscribeToInbox(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
	}
	login := r.URL.Query().Get("login")
	pass := r.URL.Query().Get("password")
	c.mediator.Handle(contract.ProxySubCmd{Login: login, Pass: pass}, r.Context())
}
