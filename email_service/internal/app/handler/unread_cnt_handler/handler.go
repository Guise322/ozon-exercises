package unread_cnt_handler

import (
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (h *unreadCntHandler) Handle(req *contract.UnreadCountRequest) (*contract.UnreadReqResult, error) {
	return h.cl.GetUnreadEmailCnt(&contract.UnreadCountRequest{
		Ctx:        req.Ctx,
		EmailLogin: req.EmailLogin,
		EmailPass:  req.EmailPass,
	})
}
