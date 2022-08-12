package app

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/handler"
	interf "github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
)

type proxyMediator struct {
	notifHandler     handler.NotifCmdHandler
	subHandler       handler.SubCmdHandler
	unreadCntHandler handler.UnreadCntHandler
}

func NewProxyMediator(cl interf.SubClient) *proxyMediator {
	return &proxyMediator{
		notifHandler:     handler.NotifCmdHandler{},
		subHandler:       handler.SubCmdHandler{SubClient: cl},
		unreadCntHandler: handler.UnreadCntHandler{},
	}
}
