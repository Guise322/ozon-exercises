package notif_client

import (
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/email_service/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type notifClient struct {
	newEmailNotifClient pb.NewEmailNotifClient
}

func NewNotifClient(confPath string) (*notifClient, error) {
	var conf NotifClientConf
	err := internal.ReadConfig(&conf, confPath)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(
		fmt.Sprintf("%v:%v", conf.NotifClient.Host, conf.NotifClient.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	grpcClient := pb.NewNewEmailNotifClient(conn)
	return &notifClient{newEmailNotifClient: grpcClient}, nil
}
