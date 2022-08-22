package unread_cnt_client

import (
	"context"
	"errors"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

type unreadCntClient struct{}

func NewUnreadCntClient() (*unreadCntClient, error) {
	return &unreadCntClient{}, nil
}

func (c *unreadCntClient) GetUnreadEmailCnt(
	ctx context.Context,
	req *contract.UnreadCountRequest,
) (*contract.UnreadReqResult, error) {
	return nil, errors.New("you're great (GetUnreadEmailCnt)")
}
