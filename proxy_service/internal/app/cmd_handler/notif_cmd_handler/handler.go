package notif_cmd_handler

import (
	"context"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			return nil, status.Error(codes.DeadlineExceeded, "processing too slow")
		default:
			return nil, status.Error(codes.Canceled, "canceled")
		}
	}
}
