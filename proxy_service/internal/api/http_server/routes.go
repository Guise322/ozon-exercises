package http_server

import "net/http"

func (s httpServer) UseRoutes() {
	http.HandleFunc("/subscribe", s.subscribeToInbox)
}
