package server

import (
	"context"
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (s server) GetUnreadEmailCount(
	ctx context.Context,
	in *pb.UnreadCountRequest,
) (*pb.UnreadCountResponse, error) {
	res, err := s.mediator.Handle(ctx, contract.UnreadCountRequest{
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
