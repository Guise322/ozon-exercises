package http_server

import (
	"net/http"

	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type httpServer struct {
	mediator mediator.Mediator
}

func newHTTPServer(med mediator.Mediator) *httpServer {
	return &httpServer{mediator: med}
}

func (s *httpServer) subscribeToInbox(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")
	pass := r.URL.Query().Get("password")
	_, err := s.mediator.Handle(&contract.ProxySubCmd{Ctx: r.Context(), Login: login, Pass: pass})
	if err != nil {
		writeServerResult(&w, http.StatusInternalServerError, err)
		return
	}
	writeServerResult(&w, http.StatusOK, nil)
}

func (s *httpServer) getUnreadEmailCnt(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")
	pass := r.URL.Query().Get("password")
	res, err := s.mediator.Handle(&contract.UnreadCntReq{
		Ctx:   r.Context(),
		Login: login,
		Pass:  pass,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeServerResult(&w, http.StatusOK, res)
}
