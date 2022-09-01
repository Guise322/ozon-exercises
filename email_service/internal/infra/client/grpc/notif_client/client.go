package notif_client

import (
	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/email_service/internal/conf"
	"github.com/Guise322/ozon-exercises/email_service/internal/infra/client/grpc/grpc_conn"
)

type NotifClient struct {
	newEmailNotifClient pb.NewEmailNotifClient
}

func NewNotifClient(c conf.NotifClientConf) (*NotifClient, error) {
	conn, err := grpc_conn.CreatGRPCConn(c.NotifClient.Host, c.NotifClient.Port)
	if err != nil {
		return nil, err
	}
	grpcClient := pb.NewNewEmailNotifClient(conn)
	return &NotifClient{newEmailNotifClient: grpcClient}, nil
}

// func (c *NotifClient) SendNotif(ctx context.Context, cmd *contract.NotifCmd) error {
// 	_, err := c.newEmailNotifClient.Notify(context.Background(), &pb.NewEmailCmd{
// 		Id:      cmd.Id,
// 		From:    cmd.From,
// 		Message: cmd.Message,
// 	})
// 	return err
// }
