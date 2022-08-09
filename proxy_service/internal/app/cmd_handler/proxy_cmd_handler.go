package cmd_handler

import (
	"context"
	"time"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type ProxyCmdHandler struct{}

func (ProxyCmdHandler) Handle(ctx context.Context, cmd contract.ProxySubCmd) {
	tiker := time.NewTicker(2000 * time.Millisecond)
	<-tiker.C
}
