package server

import (
	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/email_service/internal/app"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/interf"
)

type server struct {
	pb.UnimplementedEmailServiceServer
	mediator app.Mediator
}

func newGRPCServer(cl interf.NotifClient) *server {
	return &server{mediator: app.NewEmailMediator(cl)}
}
