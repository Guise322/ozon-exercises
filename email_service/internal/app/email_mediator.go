package app

import "github.com/Guise322/ozon-exercises/email_service/internal/app/handler"

type emailMediator struct {
	cmdHandler handler.SubCmdHandler
	reqHandler handler.EmailReqHandler
}

func NewEmailMediator() *emailMediator {
	return &emailMediator{
		cmdHandler: handler.SubCmdHandler{},
		reqHandler: handler.EmailReqHandler{},
	}
}
