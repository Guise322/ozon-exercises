package http_mid

import (
	"context"
	"net/http"
	"time"
)

func (v *HTTPTimeout) SetReqTimeout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rCtx := r.Context()
		ctx, cancel := context.WithTimeout(rCtx, time.Duration(v.value)*time.Millisecond)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
