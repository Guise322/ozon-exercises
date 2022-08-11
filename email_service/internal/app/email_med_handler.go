package app

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (m emailMediator) Handle(ctx context.Context, msg interface{}) (interface{}, error) {
	switch mes := msg.(type) {
	case contract.UnreadCountRequest:
		return m.reqHandler.Handle(ctx, mes)
	case contract.SubscribtionCmd:
		return m.cmdHandler.Handle(ctx, mes)
	default:
		return nil, errors.New("undefined msg")
	}
}
