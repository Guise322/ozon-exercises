package unread_cnt_client

type unreadCntClient struct{}

func NewUnreadCntClient() (*unreadCntClient, error) {
	return &unreadCntClient{}, nil
}
