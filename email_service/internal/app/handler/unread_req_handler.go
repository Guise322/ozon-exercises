package handler

import (
	"context"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

type EmailReqHandler struct{}

func (h *EmailReqHandler) Handle(ctx context.Context, req contract.UnreadCountRequest) (*contract.UnreadReqResult, error) {
	if req.EmailLogin == "guise322@ya.ru" {
		return &contract.UnreadReqResult{
			MessageCount: 123,
			Error:        "",
		}, nil
	}
	return &contract.UnreadReqResult{}, nil
}
