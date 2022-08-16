package http_server

import (
	"encoding/json"
	"net/http"
)

func writeServerResult(w *http.ResponseWriter, statusCode int, data interface{}) {
	(*w).WriteHeader(statusCode)
	if data != nil {
		(*w).Header().Set("Content-Type", "application/json")
		json.NewEncoder(*w).Encode(data)
	}
}
