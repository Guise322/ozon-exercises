package server

import (
	"fmt"
	"net"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"google.golang.org/grpc"
)

func RunGRPCSrv(path string) error {
	conf, err := readConfig(path)
	if err != nil {
		return err
	}
	address := fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer lis.Close()
	var opts []grpc.ServerOption
	useInterceptors(&opts, conf.Server.TimeoutInMils)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEmailServer(grpcServer, newGRPCServer())
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
