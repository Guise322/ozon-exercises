package http_server

import (
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/api"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

func (s *httpServer) subscribeToInbox(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")
	pass := r.URL.Query().Get("password")
	_, err := s.mediator.Handle(&contract.ProxySubCmd{Ctx: r.Context(), Login: login, Pass: pass})
	if err != nil {
		api.WriteServerResult(&w, http.StatusInternalServerError, err)
		return
	}
	api.WriteServerResult(&w, http.StatusOK, nil)
}