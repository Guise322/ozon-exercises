package email_handlers

import (
	"github.com/Guise322/ozon-exercises/email_bot/internal/app/contracts"
)

type EmailReqHandler struct{}

func (EmailReqHandler) Handle(req contracts.EmailRequest) contracts.EmailResponse {
	if req.Id == 322 {
		return contracts.EmailResponse{Id: req.Id, From: "test@mail.com", To: "user@mail.com", Text: "Hello there!"}
	}
	return contracts.EmailResponse{}
}
