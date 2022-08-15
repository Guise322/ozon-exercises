package subscr

import (
	"context"
	"errors"
	"time"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (h *subCmdHandler) Handle(cmd *contract.SubscribtionCmd) (*contract.SubCmdResult, error) {
	resCh := make(chan struct{})
	var (
		res *contract.SubCmdResult
		err error
	)
	go func() {
		res, err = h.cl.SubToInbox(&contract.SubscribtionCmd{
			Ctx:        cmd.Ctx,
			EmailLogin: cmd.EmailLogin,
			EmailPass:  cmd.EmailPass,
		})
		time.Sleep(500 * time.Millisecond)
		resCh <- struct{}{}
	}()
	select {
	case <-resCh:
		return res, err
	case <-cmd.Ctx.Done():
		switch cmd.Ctx.Err() {
		case context.DeadlineExceeded:
			return nil, errors.New("processing too slow")
		default:
			return nil, errors.New("canceled")
		}
	}
}
