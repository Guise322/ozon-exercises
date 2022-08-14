package unread_cnt_handler

import "github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"

func (h *unreadCntHandler) Handle(req *contract.UnreadCntReq) (*contract.UnreadCntResult, error) {
	return h.cl.GetUnreadEmailCnt(req)
}
