package grpc_server

import (
	"context"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

func (s server) Notify(ctx context.Context, in *pb.NewEmailCom) (*pb.Empty, error) {
	_, err := s.mediator.Handle(ctx, contract.NotifCmd{Id: in.Id})
	return &pb.Empty{}, err
}
