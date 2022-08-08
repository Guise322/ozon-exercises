package http

import "net/http"

func (s *HTTPServer) UseRoutes() {
	http.HandleFunc("/subscribe", s.cmdController.subscribeToInbox)
}
