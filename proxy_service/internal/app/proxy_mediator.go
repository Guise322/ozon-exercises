package app

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/cmd_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type ProxyMediator struct{}

func (ProxyMediator) Handle(msg interface{}, ctx context.Context) (interface{}, error) {
	switch s := msg.(type) {
	case contract.ProxySubCmd:
		h := &cmd_handler.ProxyCmdHandler{}
		h.Handle(ctx, s)
		return nil, nil
	default:
		return nil, errors.New("undefined msg")
	}
}
