package app

import (
	"github.com/Guise322/ozon-exercises/email_service/internal/app/handler"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/interf"
)

type emailMediator struct {
	cmdHandler handler.SubCmdHandler
	reqHandler handler.EmailReqHandler
}

func NewEmailMediator(cl interf.NotifClient) *emailMediator {
	return &emailMediator{
		cmdHandler: handler.SubCmdHandler{NotifClient: cl},
		reqHandler: handler.EmailReqHandler{},
	}
}
