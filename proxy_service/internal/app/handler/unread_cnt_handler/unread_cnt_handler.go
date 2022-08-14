package unread_cnt_handler

import "github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"

type unreadCntHandler struct {
	cl interf.UnreadCntClient
}

func NewUnreadCntHandler(cl interf.UnreadCntClient) *unreadCntHandler {
	return &unreadCntHandler{cl: cl}
}
