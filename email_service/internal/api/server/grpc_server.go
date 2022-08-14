package server

import (
	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/common/mediator"
)

type server struct {
	pb.UnimplementedEmailServiceServer
	mediator mediator.Mediator
}

func newGRPCServer(med mediator.Mediator) *server {
	return &server{mediator: med}
}