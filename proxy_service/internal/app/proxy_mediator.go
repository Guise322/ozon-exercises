package app

import "github.com/Guise322/ozon-exercises/proxy_service/internal/app/cmd_handler"

type proxyMediator struct {
	notifHandler cmd_handler.NotifCmdHandler
	subHandler   cmd_handler.SubCmdHandler
}

func NewProxyMediator() *proxyMediator {
	return &proxyMediator{
		notifHandler: cmd_handler.NotifCmdHandler{},
		subHandler:   cmd_handler.SubCmdHandler{},
	}
}
