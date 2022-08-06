package email_handler

import (
	"context"
	"errors"
	"time"

	"github.com/Guise322/ozon-exercises/email_bot/internal/app/contract"
)

type EmailComHandler struct{}

func (EmailComHandler) Handle(ctx context.Context) (*contract.EmailComResult, error) {
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	select {
	case <-ticker.C:
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}
	return &contract.EmailComResult{Data: "Hello there!"}, nil
}
