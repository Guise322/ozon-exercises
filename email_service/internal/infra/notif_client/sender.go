package notif_client

import (
	"context"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (c *notifClient) SendNotif(cmd *contract.NotifCmd) error {
	_, err := c.newEmailNotifClient.Notify(context.Background(), &pb.NewEmailCmd{
		Id:      cmd.Id,
		From:    cmd.From,
		Message: cmd.Message,
	})
	return err
}
