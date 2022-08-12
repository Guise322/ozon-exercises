package server

import (
	"fmt"
	"net"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/email_service/internal"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/interf"
	"google.golang.org/grpc"
)

func RunGRPCSrv(path string, cl interf.NotifClient) error {
	var conf ServConf
	err := internal.ReadConfig(&conf, path)
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
	pb.RegisterEmailServiceServer(grpcServer, newGRPCServer(cl))
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
