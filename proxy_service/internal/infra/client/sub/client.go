package sub

import (
	"context"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/conf"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/infra/client/grpc_conn"
)

type subClient struct {
	emailServiceClient pb.EmailServiceClient
}

func NewSubClient(c conf.SubClientConf) (*subClient, error) {
	conn, err := grpc_conn.CreatGRPCConn(c.SubClient.Host, c.SubClient.Port)
	if err != nil {
		return nil, err
	}
	grpcClient := pb.NewEmailServiceClient(conn)
	return &subClient{emailServiceClient: grpcClient}, nil
}

func (c *subClient) SubToInbox(cmd *contract.ProxySubCmd) error {
	_, err := c.emailServiceClient.SubscribeToInbox(context.Background(), &pb.SubscribtionCmd{
		EmailLogin: cmd.Login,
		EmailPass:  cmd.Pass,
	})
	return err
}
