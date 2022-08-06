package app

import (
	"errors"

	"github.com/Guise322/ozon-exercises/email_bot/internal/app/contracts"
	"github.com/Guise322/ozon-exercises/email_bot/internal/app/email_handlers"
)

type EmailMediator struct {
	Msg interface{}
}

func (m *EmailMediator) Handle() (interface{}, error) {
	switch o := m.Msg.(type) {
	case contracts.EmailRequest:
		h := &email_handlers.EmailReqHandler{Req: o}
		return h.Handle()
	case contracts.EmailCommand:
		h := email_handlers.EmailComHandler{}
		return h.Handle()
	default:
		return nil, errors.New("undefined msg")
	}
}
