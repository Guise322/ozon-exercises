package app

import "github.com/Guise322/ozon-exercises/proxy_service/internal/app/handler"

type proxyMediator struct {
	notifHandler handler.NotifCmdHandler
	subHandler   handler.SubCmdHandler
}

func NewProxyMediator() *proxyMediator {
	return &proxyMediator{
		notifHandler: handler.NotifCmdHandler{},
		subHandler:   handler.SubCmdHandler{},
	}
}
