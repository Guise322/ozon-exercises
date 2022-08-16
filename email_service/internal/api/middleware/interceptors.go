package middleware

import (
	"google.golang.org/grpc"
)

func UseInterceptors(opts *[]grpc.ServerOption, timeout int64) {
	timeoutInterc := newTimeoutInterc(timeout)
	*opts = append(*opts, grpc.UnaryInterceptor(timeoutInterc.setTimeout))
}
