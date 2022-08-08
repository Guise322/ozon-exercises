package grpc

import (
	"context"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
)

func (s *server) Notify(ctx context.Context, in *pb.NewEmailCom) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
