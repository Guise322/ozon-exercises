package server

import (
	"context"
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

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
