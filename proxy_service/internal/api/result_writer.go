package api

import (
	"encoding/json"
	"net/http"
)

func WriteServerResult(w *http.ResponseWriter, statusCode int, data interface{}) {
	(*w).WriteHeader(statusCode)
	if data != nil {
		(*w).Header().Set("Content-Type", "application/json")
		if d, ok := data.(error); ok {
			json.NewEncoder(*w).Encode(d.Error())
			return
		}
		json.NewEncoder(*w).Encode(data)
	}
}
