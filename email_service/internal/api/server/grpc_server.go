package server

import (
	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"github.com/Guise322/ozon-exercises/email_service/internal/app"
)

type server struct {
	pb.UnimplementedEmailServer
	mediator app.Mediator
}

func newGRPCServer() *server {
	return &server{mediator: app.NewEmailMediator()}
}
