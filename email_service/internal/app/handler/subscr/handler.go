package subscr

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/interf"
)

type subCmdHandler struct {
	cl interf.SubClient
}

func NewSubCmdHandler(cl interf.SubClient) *subCmdHandler {
	return &subCmdHandler{cl: cl}
}

func (h *subCmdHandler) Handle(ctx context.Context, cmd *contract.SubscribtionCmd) (interface{}, error) {
	resCh := make(chan struct{})
	var err error
	go func() {
		err = h.cl.SubToInbox(ctx, &contract.SubscribtionCmd{
			EmailLogin: cmd.EmailLogin,
			EmailPass:  cmd.EmailPass,
		})
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
