package api

import (
	"context"
	"fmt"
	"net"

	"github.com/Guise322/ozon-exercises/common/email_service_pb"
	"github.com/Guise322/ozon-exercises/email_service/internal/api/middleware"
	"github.com/Guise322/ozon-exercises/email_service/internal/app"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"

	"google.golang.org/grpc"
)

type server struct {
	email_service_pb.UnimplementedEmailServer
}

func RunGRPCSrv(timeout int64, lis net.Listener) error {
	timeout_interc := middleware.TimeoutInterceptor{Timeout: timeout}
	id_valid_interc := middleware.IdValidInterceptor{}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.ChainUnaryInterceptor(id_valid_interc.ValidateId, timeout_interc.SetTimeout))
	grpcServer := grpc.NewServer(opts...)
	email_service_pb.RegisterEmailServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *server) GetEmail(
	ctx context.Context,
	in *email_service_pb.EmailRequest) (*email_service_pb.EmailResponse, error) {
	med := app.EmailMediator{Msg: contract.EmailRequest{Id: in.Id}}
	res, err := med.Handle(ctx)
	if err != nil {
		return nil, err
	}
	reqRes, ok := res.(*contract.EmailReqResult)
	if !ok {
		return nil, fmt.Errorf("wrong response format: %v", reqRes)
	}
	return &email_service_pb.EmailResponse{Id: reqRes.Id, From: reqRes.From, To: reqRes.To, Text: reqRes.Text}, nil
}

func (s *server) SubscribeToInbox(
	ctx context.Context,
	in *email_service_pb.SubscribeCom) (*email_service_pb.ComResponse, error) {
	med := app.EmailMediator{Msg: contract.EmailCommand{}}
	res, err := med.Handle(ctx)
	if err != nil {
		return nil, err
	}
	reqRes, ok := res.(*contract.EmailComResult)
	if !ok {
		return nil, fmt.Errorf("wrong response format: %v", reqRes)
	}
	return &email_service_pb.ComResponse{Result: reqRes.Data}, nil
}
