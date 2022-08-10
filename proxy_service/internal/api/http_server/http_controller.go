package http_server

import (
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

func (s httpServer) subscribeToInbox(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
	}
	login := r.URL.Query().Get("login")
	pass := r.URL.Query().Get("password")
	s.mediator.Handle(r.Context(), contract.ProxySubCmd{Login: login, Pass: pass})
}
