package server

import (
	"context"
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
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
) (*pb.SubCmdResponse, error) {
	res, err := s.mediator.Handle(&contract.SubscribtionCmd{
		Ctx:        ctx,
		EmailLogin: in.EmailLogin,
		EmailPass:  in.EmailPass,
	})
	if err != nil {
		return nil, err
	}
	cmdRes, ok := res.(*contract.SubCmdResult)
	if !ok {
		return nil, fmt.Errorf("wrong response format: %v", cmdRes)
	}
	return &pb.SubCmdResponse{Error: cmdRes.Error}, nil
}

func (s *server) GetUnreadEmailCount(
	ctx context.Context,
	in *pb.UnreadCountRequest,
) (*pb.UnreadCountResponse, error) {
	res, err := s.mediator.Handle(&contract.UnreadCountRequest{
		Ctx:        ctx,
		EmailLogin: in.EmailLogin,
		EmailPass:  in.EmailPass,
	})
	if err != nil {
		return nil, err
	}
	reqRes, ok := res.(*contract.UnreadReqResult)
	if !ok {
		return nil, fmt.Errorf("wrong response format: %v", reqRes)
	}
	return &pb.UnreadCountResponse{MessageCount: reqRes.MessageCount, Error: reqRes.Error}, nil
}
