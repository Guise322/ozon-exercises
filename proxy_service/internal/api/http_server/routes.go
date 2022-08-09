package http_server

import "net/http"

func (s *HTTPServer) UseRoutes() {
	http.HandleFunc("/subscribe", s.subscribeToInbox)
}
