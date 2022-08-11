package api

import (
	"encoding/json"
	"net/http"
)

func WriteServerResult(w *http.ResponseWriter, statusCode int, data interface{}) {
	(*w).WriteHeader(statusCode)
	if data != nil {
		(*w).Header().Set("Content-Type", "application/json")
		json.NewEncoder(*w).Encode(data)
	}
}
