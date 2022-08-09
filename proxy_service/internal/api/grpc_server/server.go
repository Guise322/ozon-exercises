package grpc_server

import (
	"context"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
)

type server struct {
	pb.UnimplementedNewEmailNotifServer
}

func (s *server) Notify(ctx context.Context, in *pb.NewEmailCom) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
