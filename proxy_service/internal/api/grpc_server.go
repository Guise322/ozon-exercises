package api

import (
	"context"
	"net"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/middleware"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNewEmailNotifServer
}

func RunGRPCSrv(lis net.Listener) error {
	id_valid_interc := middleware.IdValidInterceptor{}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(id_valid_interc.ValidateId))
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterNewEmailNotifServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *server) Notify(ctx context.Context, in *pb.NewEmailCom) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
