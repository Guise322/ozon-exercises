package middleware

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type TimeoutInterceptor struct {
	Timeout int64
}

func (t TimeoutInterceptor) SetTimeout(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	timeCtx, cancel := context.WithTimeout(ctx, time.Duration(t.Timeout)*time.Millisecond)
	defer cancel()
	return handler(timeCtx, req)
}
