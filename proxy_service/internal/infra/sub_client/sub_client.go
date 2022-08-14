package sub_client

import (
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/proxy_service/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type subClient struct {
	emailServiceClient pb.EmailServiceClient
}

func NewSubClient(confPath string) (*subClient, error) {
	var conf SubClientConf
	err := internal.ReadConfig(&conf, confPath)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", conf.SubClient.Host, conf.SubClient.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	grpcClient := pb.NewEmailServiceClient(conn)
	return &subClient{emailServiceClient: grpcClient}, nil
}
