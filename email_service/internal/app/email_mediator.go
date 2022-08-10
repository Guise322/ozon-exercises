package app

import (
	"github.com/Guise322/ozon-exercises/email_service/internal/app/handler/cmd_handler"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/handler/req_handler"
)

type emailMediator struct {
	cmdHandler cmd_handler.EmailCmdHandler
	reqHandler req_handler.EmailReqHandler
}

func NewEmailMediator() *emailMediator {
	return &emailMediator{
		cmdHandler: cmd_handler.EmailCmdHandler{},
		reqHandler: req_handler.EmailReqHandler{},
	}
}
