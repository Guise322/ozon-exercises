package grpc_server

import (
	"net"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNewEmailNotifServer
}

func RunGRPCSrv(lis net.Listener) error {

	var opts []grpc.ServerOption
	UseInterceptors(&opts)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterNewEmailNotifServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
