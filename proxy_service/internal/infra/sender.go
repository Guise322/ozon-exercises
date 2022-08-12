package infra

import (
	"context"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

func (c *subClient) SubToInbox(cmd *contract.ProxySubCmd) error {
	_, err := c.emailServiceClient.SubscribeToInbox(context.Background(), &pb.SubscribtionCmd{
		EmailLogin: cmd.Login,
		EmailPass:  cmd.Pass,
	})
	return err
}
