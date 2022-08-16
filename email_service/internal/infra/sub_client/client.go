package sub_client

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

type subClient struct{}

func NewSubClient() (*subClient, error) {
	return &subClient{}, nil
}

func (c *subClient) SubToInbox(ctx context.Context, cmd *contract.SubscribtionCmd) error {
	return errors.New("you're great (SubToInbox)")
}
