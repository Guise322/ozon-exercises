package api

import (
	"context"
	"errors"
	"net"

	pb "github.com/Guise322/ozon-exercises/email_bot/proto"

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
	if in.GetId() == 322 {
		return &pb.EmailResponse{Id: 322, From: "test@mail.com", To: "user@mail.com", Text: "Hello there!"}, nil
	}
	return nil, errors.New("id is not available")
}
