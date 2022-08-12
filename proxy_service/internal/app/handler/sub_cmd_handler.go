package handler

import (
	"context"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	interf "github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
)

type SubCmdHandler struct {
	SubClient interf.SubClient
}

func (h *SubCmdHandler) Handle(ctx context.Context, cmd *contract.ProxySubCmd) error {
	return h.SubClient.SubToInbox(cmd)
}
