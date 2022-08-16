package grpc_server

import (
	"fmt"
	"net"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/conf"
	"google.golang.org/grpc"
)

func RunGRPCSrv(c conf.GRPCConf, med mediator.Mediator) error {

	address := fmt.Sprintf("%v:%v", c.Server.Host, c.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer lis.Close()
	var opts []grpc.ServerOption
	useInterceptors(&opts, c.Server.Timeout)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterNewEmailNotifServer(grpcServer, newGRPCServer(med))
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
