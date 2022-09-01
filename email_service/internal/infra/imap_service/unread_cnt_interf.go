package imap_service

type UnreadCntGetter interface {
	Connect(imapAddr, email, password string) error
	Disconnect() error
	GetUnreadMessageCount() (int, error)
}
