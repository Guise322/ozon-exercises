package sub_cmd_handler

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
)

type subCmdHandler struct {
	subClient interf.SubClient
}

func NewSubCmdHandler(subClient interf.SubClient) *subCmdHandler {
	return &subCmdHandler{subClient: subClient}
}

func (h *subCmdHandler) Handle(ctx context.Context, cmd *contract.ProxySubCmd) (interface{}, error) {
	resCh := make(chan struct{})
	var err error
	go func() {
		err = h.subClient.SubToInbox(ctx, cmd)
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
