package email_handler

import (
	"github.com/Guise322/ozon-exercises/email_bot/internal/app/contract"
)

type EmailReqHandler struct {
	Req contract.EmailRequest
}

func (h *EmailReqHandler) Handle() (*contract.EmailReqResult, error) {
	if h.Req.Id == 322 {
		return &contract.EmailReqResult{
			Id:   h.Req.Id,
			From: "test@mail.com",
			To:   "user@mail.com",
			Text: "Hello there!"}, nil
	}
	return &contract.EmailReqResult{}, nil
}
