package sub_cmd_handler

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *subCmdHandler) Handle(cmd *contract.ProxySubCmd) (interface{}, error) {
	resCh := make(chan struct{})
	var err error
	go func() {
		err = h.subClient.SubToInbox(cmd)
		resCh <- struct{}{}
	}()

	select {
	case <-resCh:
		return nil, err
	case <-cmd.Ctx.Done():
		switch cmd.Ctx.Err() {
		case context.DeadlineExceeded:
			return empty.Empty{}, errors.New("processing too slow")
		default:
			return empty.Empty{}, errors.New("canceled")
		}
	}
}
