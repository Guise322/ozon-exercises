package grpc_server

import (
	"fmt"
	"net"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"google.golang.org/grpc"
)

func RunGRPCSrv(confPath string) error {
	conf, err := readConfig(confPath)
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
	useInterceptors(&opts)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterNewEmailNotifServer(grpcServer, newGRPCServer())
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
