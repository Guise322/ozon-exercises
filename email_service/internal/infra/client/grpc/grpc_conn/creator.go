package grpc_conn

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreatGRPCConn(host string, port int64) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return conn, err
}
