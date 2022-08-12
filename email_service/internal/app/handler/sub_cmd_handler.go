package handler

import (
	"context"
	"errors"
	"time"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/interf"
)

type SubCmdHandler struct {
	NotifClient interf.NotifClient
}

func (h *SubCmdHandler) Handle(ctx context.Context, cmd contract.SubscribtionCmd) (*contract.SubCmdResult, error) {
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	select {
	case <-ticker.C:
		h.NotifClient.SendNotif(&contract.NotifCmd{Id: 322, From: "your friend", Message: "You are great!"})
	case <-ctx.Done():
		return nil, errors.New("timeout")
	}
	return &contract.SubCmdResult{Error: "Hello there!"}, nil
}
