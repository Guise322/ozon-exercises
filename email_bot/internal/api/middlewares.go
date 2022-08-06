package api

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type timeoutInterceptor struct {
	timeout int64
}

func (t timeoutInterceptor) SetTimeout(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	timeCtx, cancel := context.WithTimeout(ctx, time.Duration(t.timeout)*time.Millisecond)
	defer cancel()
	return handler(timeCtx, req)
}
