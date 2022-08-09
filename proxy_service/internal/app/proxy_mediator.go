package app

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/cmd_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type ProxyMediator struct{}

func (ProxyMediator) Handle(ctx context.Context, msg interface{}) (interface{}, error) {
	switch s := msg.(type) {
	case contract.ProxySubCmd:
		h := cmd_handler.SubCmdHandler{}
		return nil, h.Handle(ctx, s)
	case contract.NotifCmd:
		h := cmd_handler.NotifCmdHandler{}
		h.Handle(ctx, s)
		return nil, h.Handle(ctx, s)
	default:
		return nil, errors.New("undefined msg")
	}
}
