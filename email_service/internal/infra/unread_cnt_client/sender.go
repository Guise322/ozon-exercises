package unread_cnt_client

import (
	"errors"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (c *unreadCntClient) GetUnreadEmailCnt(req *contract.UnreadCountRequest) (*contract.UnreadReqResult, error) {
	return nil, errors.New("you're great (GetUnreadEmailCnt)")
}
