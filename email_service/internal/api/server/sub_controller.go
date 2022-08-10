package server

import (
	"context"
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (s server) SubscribeToInbox(
	ctx context.Context,
	in *pb.SubscribeCom) (*pb.ComResponse, error) {
	res, err := s.mediator.Handle(ctx, contract.EmailCommand{})
	if err != nil {
		return nil, err
	}
	reqRes, ok := res.(*contract.EmailCmdResult)
	if !ok {
		return nil, fmt.Errorf("wrong response format: %v", reqRes)
	}
	return &pb.ComResponse{Result: reqRes.Data}, nil
}
