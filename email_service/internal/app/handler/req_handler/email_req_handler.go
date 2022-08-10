package req_handler

import (
	"context"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

type EmailReqHandler struct{}

func (h *EmailReqHandler) Handle(ctx context.Context, req contract.EmailRequest) (*contract.EmailReqResult, error) {
	if req.Id == 322 {
		return &contract.EmailReqResult{
			Id:   req.Id,
			From: "test@mail.com",
			To:   "user@mail.com",
			Text: "Hello there!"}, nil
	}
	return &contract.EmailReqResult{}, nil
}
