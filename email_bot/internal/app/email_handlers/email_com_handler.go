package email_handlers

import "github.com/Guise322/ozon-exercises/email_bot/internal/app/contracts"

type EmailComHandler struct{}

func (EmailComHandler) Handle() (*contracts.EmailComResult, error) {
	return &contracts.EmailComResult{Data: "Hello there!"}, nil
}
