package interf

import (
	"context"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type NotifClient interface {
	Notify(ctx context.Context, cmd *contract.NotifCmd) error
}
