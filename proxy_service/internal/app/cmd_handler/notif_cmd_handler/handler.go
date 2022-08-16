package notif_cmd_handler

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
	"github.com/golang/protobuf/ptypes/empty"
)

type notifCmdHandler struct {
	cl interf.NotifClient
}

func NewNotifCmdHandler(cl interf.NotifClient) *notifCmdHandler {
	return &notifCmdHandler{cl: cl}
}

func (h *notifCmdHandler) Handle(cmd *contract.NotifCmd) (empty.Empty, error) {
	resCh := make(chan struct{})
	var err error
	go func() {
		_, err = h.cl.Notify(cmd)
		resCh <- struct{}{}
	}()
	select {
	case <-resCh:
		return empty.Empty{}, err
	case <-cmd.Ctx.Done():
		switch cmd.Ctx.Err() {
		case context.DeadlineExceeded:
			return empty.Empty{}, errors.New("processing too slow")
		default:
			return empty.Empty{}, errors.New("canceled")
		}
	}
}
