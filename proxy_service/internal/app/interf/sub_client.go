package interf

import (
	"context"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type SubClient interface {
	SubToInbox(ctx context.Context, cmd *contract.ProxySubCmd) error
}
