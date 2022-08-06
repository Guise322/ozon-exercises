package app

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/email_bot/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_bot/internal/app/email_handler"
)

type EmailMediator struct {
	Msg interface{}
}

func (m *EmailMediator) Handle(ctx context.Context) (interface{}, error) {
	switch o := m.Msg.(type) {
	case contract.EmailRequest:
		h := &email_handler.EmailReqHandler{Req: o}
		return h.Handle(ctx)
	case contract.EmailCommand:
		h := email_handler.EmailComHandler{}
		return h.Handle(ctx)
	default:
		return nil, errors.New("undefined msg")
	}
}
