package handler

import (
	"errors"
	"time"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/interf"
)

type SubCmdHandler struct {
	NotifClient interf.NotifClient
}

func (h *SubCmdHandler) Handle(cmd contract.SubscribtionCmd) (*contract.SubCmdResult, error) {
	ticker := time.NewTicker(1 * time.Nanosecond)
	defer ticker.Stop()
	select {
	case <-ticker.C:
		h.NotifClient.SendNotif(&contract.NotifCmd{Id: 322, From: "your friend", Message: "You are great!"})
	case <-cmd.Ctx.Done():
		return nil, errors.New("timeout")
	}
	return &contract.SubCmdResult{Error: "Hello there!"}, errors.New("Ho-ho")
}
