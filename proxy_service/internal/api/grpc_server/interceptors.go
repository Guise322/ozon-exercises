package grpc_server

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/middleware"
	"google.golang.org/grpc"
)

func UseInterceptors(opts *[]grpc.ServerOption) {
	id_valid_interc := middleware.IdValidInterceptor{}
	*opts = append(*opts, grpc.UnaryInterceptor(id_valid_interc.ValidateId))
}
