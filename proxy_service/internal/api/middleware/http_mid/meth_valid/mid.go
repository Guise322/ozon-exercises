package meth_valid

import "net/http"

type HTTPMethValid struct {
	method string
}

func NewHTTPValidator(method string) *HTTPMethValid {
	return &HTTPMethValid{method: method}
}

func (h *HTTPMethValid) ValidateHTTPMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != h.method {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		next.ServeHTTP(w, r)
	})
}
