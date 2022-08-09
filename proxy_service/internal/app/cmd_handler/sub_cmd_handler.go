package cmd_handler

import (
	"context"
	"time"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type SubCmdHandler struct{}

func (SubCmdHandler) Handle(ctx context.Context, cmd contract.ProxySubCmd) error {
	tiker := time.NewTicker(2000 * time.Millisecond)
	<-tiker.C
	return nil
}
