package app

import (
	"errors"

	"github.com/Guise322/ozon-exercises/email_bot/internal/app/contracts"
	"github.com/Guise322/ozon-exercises/email_bot/internal/app/email_handlers"
)

type EmailMediator struct{}

func (EmailMediator) Handle(msg interface{}) (interface{}, error) {
	switch o := msg.(type) {
	case contracts.EmailRequest:
		var h contracts.EmailReqHandler
		h = email_handlers.EmailReqHandler{}
		return h.Handle(o), nil
	case contracts.EmailCommand:
		return nil, errors.New("command is undefined")
	default:
		return nil, errors.New("undefined msg")
	}
}
