package app

import (
	"errors"

	"github.com/Guise322/ozon-exercises/email_bot/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_bot/internal/app/email_handler"
)

type EmailMediator struct {
	Msg interface{}
}

func (m *EmailMediator) Handle() (interface{}, error) {
	switch o := m.Msg.(type) {
	case contract.EmailRequest:
		h := &email_handler.EmailReqHandler{Req: o}
		return h.Handle()
	case contract.EmailCommand:
		h := email_handler.EmailComHandler{}
		return h.Handle()
	default:
		return nil, errors.New("undefined msg")
	}
}
