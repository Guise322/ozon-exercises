package sub_client

import (
	"errors"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (c *subClient) SubToInbox(cmd *contract.SubscribtionCmd) (*contract.SubCmdResult, error) {
	return nil, errors.New("you're great (SubToInbox)")
}
