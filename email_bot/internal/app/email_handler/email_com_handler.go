package email_handler

import "github.com/Guise322/ozon-exercises/email_bot/internal/app/contract"

type EmailComHandler struct{}

func (EmailComHandler) Handle() (*contract.EmailComResult, error) {
	return &contract.EmailComResult{Data: "Hello there!"}, nil
}
