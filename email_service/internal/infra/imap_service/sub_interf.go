package imap_service

import "github.com/emersion/go-imap"

type Subscriber interface {
	Connect(imapAddr, email, password string) error
	Disconnect() error
	SubToInbox(mesCh chan<- *imap.Message) error
	UnsubFromInbox() error
}
