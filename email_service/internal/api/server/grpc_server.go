package server

import (
	"context"
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"github.com/Guise322/ozon-exercises/email_service/internal/app"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

type server struct {
	pb.UnimplementedEmailServer
}

func (s *server) GetEmail(
	ctx context.Context,
	in *pb.EmailRequest) (*pb.EmailResponse, error) {
	med := app.EmailMediator{Msg: contract.EmailRequest{Id: in.Id}}
	res, err := med.Handle(ctx)
	if err != nil {
		return nil, err
	}
	reqRes, ok := res.(*contract.EmailReqResult)
	if !ok {
		return nil, fmt.Errorf("wrong response format: %v", reqRes)
	}
	return &pb.EmailResponse{Id: reqRes.Id, From: reqRes.From, To: reqRes.To, Text: reqRes.Text}, nil
}

func (s *server) SubscribeToInbox(
	ctx context.Context,
	in *pb.SubscribeCom) (*pb.ComResponse, error) {
	med := app.EmailMediator{Msg: contract.EmailCommand{}}
	res, err := med.Handle(ctx)
	if err != nil {
		return nil, err
	}
	reqRes, ok := res.(*contract.EmailCmdResult)
	if !ok {
		return nil, fmt.Errorf("wrong response format: %v", reqRes)
	}
	return &pb.ComResponse{Result: reqRes.Data}, nil
}
