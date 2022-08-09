package cmd_handler

import (
	"context"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type NotifCmdHandler struct{}

func (NotifCmdHandler) Handle(ctx context.Context, cmd contract.NotifCmd) error {
	return nil
}
