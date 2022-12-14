package grpc_server

import (
	"context"
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedNewEmailNotifServer
	mediator mediator.Mediator
}

func newGRPCServer(med mediator.Mediator) *server {
	return &server{mediator: med}
}

func (s server) Notify(ctx context.Context, in *pb.NewEmailCmd) (*empty.Empty, error) {
	_, err := s.mediator.Handle(ctx, &contract.NotifCmd{
		Id:      in.Id,
		From:    in.From,
		Message: in.Message,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("%v (proxy_service)", err.Error()))
	}
	return &empty.Empty{}, nil
}
