package server

import (
	"fmt"
	"net"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/email_service/internal/conf"
	"google.golang.org/grpc"
)

func RunGRPCSrv(c conf.ServConf, med mediator.Mediator) error {
	address := fmt.Sprintf("%v:%v", c.Server.Host, c.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer lis.Close()
	var opts []grpc.ServerOption
	useInterceptors(&opts, c.Server.TimeoutInMilSec)
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEmailServiceServer(grpcServer, newGRPCServer(med))
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
