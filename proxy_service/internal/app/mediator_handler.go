package app

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

func (m *proxyMediator) Handle(ctx context.Context, msg interface{}) (interface{}, error) {
	switch s := msg.(type) {
	case contract.ProxySubCmd:
		return nil, m.subHandler.Handle(ctx, &s)
	case contract.NotifCmd:
		return nil, m.notifHandler.Handle(ctx, s)
	case contract.UnreadCntReq:
		return m.unreadCntHandler.Handle(ctx, s)
	default:
		return nil, errors.New("undefined msg")
	}
}
