package grpc_server

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/middleware/grpc_mid"
	"google.golang.org/grpc"
)

func useInterceptors(opts *[]grpc.ServerOption, timeout int64) {
	timeoutInterc := grpc_mid.NewTimeoutInterc(timeout)
	*opts = append(*opts, grpc.UnaryInterceptor(timeoutInterc.SetTimeout))
}
