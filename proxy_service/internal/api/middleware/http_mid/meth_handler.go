package http_mid

import "net/http"

func (h *HTTPMethValid) ValidateHTTPMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != h.method {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		next.ServeHTTP(w, r)
	})
}
