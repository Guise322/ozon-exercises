package server

import (
	"github.com/Guise322/ozon-exercises/email_service/internal/api/middleware"
	"google.golang.org/grpc"
)

func useInterceptors(opts *[]grpc.ServerOption, timeout int64) {
	timeoutInterc := middleware.TimeoutInterceptor{Timeout: timeout}
	*opts = append(*opts, grpc.UnaryInterceptor(timeoutInterc.SetTimeout))
}
