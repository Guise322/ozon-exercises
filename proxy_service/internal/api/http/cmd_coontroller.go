package http

import "net/http"

type cmdController struct{}

func (cmdController) subscribeToInbox(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
