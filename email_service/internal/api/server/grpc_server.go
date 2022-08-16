package server

import (
	"context"
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedEmailServiceServer
	mediator mediator.Mediator
}

func newGRPCServer(med mediator.Mediator) *server {
	return &server{mediator: med}
}

func (s *server) SubscribeToInbox(
	ctx context.Context,
	in *pb.SubscribtionCmd,
) (*empty.Empty, error) {
	_, err := s.mediator.Handle(ctx, &contract.SubscribtionCmd{
		EmailLogin: in.EmailLogin,
		EmailPass:  in.EmailPass,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("email_service: %v", err.Error()))
	}
	return &empty.Empty{}, nil
}

func (s *server) GetUnreadEmailCount(
	ctx context.Context,
	in *pb.UnreadCountRequest,
) (*pb.UnreadCountResponse, error) {
	res, err := s.mediator.Handle(ctx, &contract.UnreadCountRequest{
		EmailLogin: in.EmailLogin,
		EmailPass:  in.EmailPass,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("email_service: %v", err.Error()))
	}
	reqRes, ok := res.(*contract.UnreadReqResult)
	if !ok {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("email_service: wrong response format: %v", reqRes),
		)
	}
	return &pb.UnreadCountResponse{MessageCount: reqRes.MessageCount}, nil
}
