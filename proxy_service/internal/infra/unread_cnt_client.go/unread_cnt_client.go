package unread_cnt_client

import (
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/proxy_service/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type unreadCntClient struct {
	emailServiceClient pb.EmailServiceClient
}

func NewUnreadCntClient(confPath string) (*unreadCntClient, error) {
	var conf UnreadCntClientConf
	err := internal.ReadConfig(&conf, confPath)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", conf.UnreadCntClient.Host, conf.UnreadCntClient.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	grpcClient := pb.NewEmailServiceClient(conn)
	return &unreadCntClient{emailServiceClient: grpcClient}, nil
}
