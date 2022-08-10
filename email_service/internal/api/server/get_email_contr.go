package server

import (
	"context"
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (s server) GetEmail(
	ctx context.Context,
	in *pb.EmailRequest) (*pb.EmailResponse, error) {
	res, err := s.mediator.Handle(ctx, contract.EmailRequest{Id: in.Id})
	if err != nil {
		return nil, err
	}
	reqRes, ok := res.(*contract.EmailReqResult)
	if !ok {
		return nil, fmt.Errorf("wrong response format: %v", reqRes)
	}
	return &pb.EmailResponse{Id: reqRes.Id, From: reqRes.From, To: reqRes.To, Text: reqRes.Text}, nil
}
