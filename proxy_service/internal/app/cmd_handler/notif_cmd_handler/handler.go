package notif_cmd_handler

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
)

type notifCmdHandler struct {
	cl interf.NotifClient
}

func NewNotifCmdHandler(cl interf.NotifClient) *notifCmdHandler {
	return &notifCmdHandler{cl: cl}
}

func (h *notifCmdHandler) Handle(ctx context.Context, cmd *contract.NotifCmd) (interface{}, error) {
	resCh := make(chan struct{})
	var err error
	go func() {
		err = h.cl.Notify(ctx, cmd)
		resCh <- struct{}{}
	}()
	select {
	case <-resCh:
		return nil, err
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			return nil, errors.New("processing too slow")
		default:
			return nil, errors.New("canceled")
		}
	}
}
