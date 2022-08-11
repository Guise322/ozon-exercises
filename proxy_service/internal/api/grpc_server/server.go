package grpc_server

import (
	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app"
)

type server struct {
	pb.UnimplementedNewEmailNotifServer
	mediator app.Mediator
}

func newGRPCServer() *server {
	return &server{mediator: app.NewProxyMediator()}
}
