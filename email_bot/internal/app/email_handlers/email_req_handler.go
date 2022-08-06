package email_handlers

import (
	"github.com/Guise322/ozon-exercises/email_bot/internal/app/contracts"
)

type EmailReqHandler struct {
	Req contracts.EmailRequest
}

func (h *EmailReqHandler) Handle() (*contracts.EmailReqResult, error) {
	if h.Req.Id == 322 {
		return &contracts.EmailReqResult{
			Id:   h.Req.Id,
			From: "test@mail.com",
			To:   "user@mail.com",
			Text: "Hello there!"}, nil
	}
	return &contracts.EmailReqResult{}, nil
}
