package cmd_handler

import (
	"context"
	"errors"
	"time"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

type EmailCmdHandler struct{}

func (EmailCmdHandler) Handle(ctx context.Context) (*contract.EmailCmdResult, error) {
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	select {
	case <-ticker.C:
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}
	return &contract.EmailCmdResult{Data: "Hello there!"}, nil
}
