package api

import (
	"context"
	"fmt"
	"net"

	"github.com/Guise322/ozon-exercises/email_bot/internal/app"
	"github.com/Guise322/ozon-exercises/email_bot/internal/app/contracts"
	pb "github.com/Guise322/ozon-exercises/email_bot/protos"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEmailServer
}

func NewServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterEmailServer(grpcServer, &server{})
	return grpcServer
}

func Start(server *grpc.Server, lis net.Listener) error {
	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *server) GetEmail(ctx context.Context, in *pb.EmailRequest) (*pb.EmailResponse, error) {
	med := app.EmailMediator{}
	res, err := med.Handle(contracts.EmailRequest{Id: in.Id})
	if err != nil {
		return nil, err
	}
	reqRes, ok := res.(contracts.EmailResponse)
	if !ok {
		return nil, fmt.Errorf("wrong response: %v", reqRes)
	}
	return &pb.EmailResponse{Id: reqRes.Id, From: reqRes.From, To: reqRes.To, Text: reqRes.Text}, nil
}
