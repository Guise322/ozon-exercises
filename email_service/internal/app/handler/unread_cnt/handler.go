package unread_cnt

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/interf"
)

type unreadCntHandler struct {
	cl interf.UnreadCntClient
}

func NewUnreadCntHandler(cl interf.UnreadCntClient) *unreadCntHandler {
	return &unreadCntHandler{cl: cl}
}

func (h *unreadCntHandler) Handle(ctx context.Context, req *contract.UnreadCountRequest) (*contract.UnreadReqResult, error) {
	resCh := make(chan struct{})
	var (
		res *contract.UnreadReqResult
		err error
	)
	go func() {
		res, err = h.cl.GetUnreadEmailCnt(ctx, &contract.UnreadCountRequest{
			EmailLogin: req.EmailLogin,
			EmailPass:  req.EmailPass,
		})
		resCh <- struct{}{}
	}()
	select {
	case <-resCh:
		return res, err
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			return nil, errors.New("processing too slow")
		default:
			return nil, errors.New("canceled")
		}
	}
}
