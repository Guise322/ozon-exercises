package interf

import (
	"context"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

type NotifClient interface {
	SendNotif(ctx context.Context, cmd *contract.NotifCmd) error
}
