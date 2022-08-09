package server

import (
	"net"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"google.golang.org/grpc"
)

func RunGRPCSrv(timeout int64, lis net.Listener) error {
	var opts []grpc.ServerOption
	UseInterceptors(&opts, timeout)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEmailServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
