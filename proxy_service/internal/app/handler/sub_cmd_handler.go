package handler

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
)

type SubCmdHandler struct {
	SubClient interf.SubClient
}

func (h *SubCmdHandler) Handle(cmd *contract.ProxySubCmd) (interface{}, error) {
	return nil, h.SubClient.SubToInbox(cmd)
}
