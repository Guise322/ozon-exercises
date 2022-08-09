package server

import (
	"github.com/Guise322/ozon-exercises/email_service/internal/api/middleware"
	"google.golang.org/grpc"
)

func UseInterceptors(opts *[]grpc.ServerOption, timeout int64) {
	idIalidInterc := middleware.IdValidInterceptor{}
	timeoutInterc := middleware.TimeoutInterceptor{Timeout: timeout}
	*opts = append(*opts, grpc.ChainUnaryInterceptor(idIalidInterc.ValidateId, timeoutInterc.SetTimeout))
}
