package grpc_server

import (
	"context"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s server) Notify(ctx context.Context, in *pb.NewEmailCmd) (*emptypb.Empty, error) {
	_, err := s.mediator.Handle(ctx, contract.NotifCmd{Id: in.Id})
	return &emptypb.Empty{}, err
}
