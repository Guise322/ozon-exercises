package grpc_server

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/middleware"
	"google.golang.org/grpc"
)

func UseInterceptors(opts *[]grpc.ServerOption) {
	idIalidInterc := middleware.IdValidInterceptor{}
	*opts = append(*opts, grpc.UnaryInterceptor(idIalidInterc.ValidateId))
}
