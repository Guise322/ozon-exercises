package req_handler

import (
	"context"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

type EmailReqHandler struct {
	Req contract.EmailRequest
}

func (h *EmailReqHandler) Handle(ctx context.Context) (*contract.EmailReqResult, error) {
	if h.Req.Id == 322 {
		return &contract.EmailReqResult{
			Id:   h.Req.Id,
			From: "test@mail.com",
			To:   "user@mail.com",
			Text: "Hello there!"}, nil
	}
	return &contract.EmailReqResult{}, nil
}
