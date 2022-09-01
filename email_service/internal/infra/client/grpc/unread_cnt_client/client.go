package unread_cnt_client

import (
	"context"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_service/internal/infra/imap_service"
)

type unreadCntClient struct {
	unreadCntGetter imap_service.UnreadCntGetter
	imapAddr        string
}

func NewUnreadCntClient(
	unreadCntGetter imap_service.UnreadCntGetter,
	imapAddr string,
) (*unreadCntClient, error) {
	return &unreadCntClient{unreadCntGetter: unreadCntGetter, imapAddr: imapAddr}, nil
}

func (c *unreadCntClient) GetUnreadEmailCnt(
	ctx context.Context,
	req *contract.UnreadCountRequest,
) (*contract.UnreadReqResult, error) {
	err := c.unreadCntGetter.Connect(c.imapAddr, req.EmailLogin, req.EmailPass)
	if err != nil {
		return nil, err
	}
	cnt, err := c.unreadCntGetter.GetUnreadMessageCount()
	if err != nil {
		return nil, err
	}
	c.unreadCntGetter.Disconnect()
	return &contract.UnreadReqResult{MessageCount: cnt}, nil
}
