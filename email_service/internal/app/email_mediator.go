package app

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/handler/cmd_handler"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/handler/req_handler"
)

type EmailMediator struct {
}

func (EmailMediator) Handle(ctx context.Context, msg interface{}) (interface{}, error) {
	switch o := msg.(type) {
	case contract.EmailRequest:
		h := req_handler.EmailReqHandler{Req: o}
		return h.Handle(ctx)
	case contract.EmailCommand:
		h := cmd_handler.EmailCmdHandler{}
		return h.Handle(ctx)
	default:
		return nil, errors.New("undefined msg")
	}
}
