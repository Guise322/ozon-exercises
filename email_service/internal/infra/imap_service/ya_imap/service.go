package ya_imap

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type yaService struct {
	client    *client.Client
	localDone chan struct{}
	mesCh     chan<- *imap.Message
	unsubCh   chan struct{}
}

func NewYaService(mesCh chan<- *imap.Message) *yaService {
	return &yaService{mesCh: mesCh, unsubCh: make(chan struct{})}
}

func (s *yaService) Connect(imapAddr, email, password string) error {
	var err error
	s.client, err = client.DialTLS(imapAddr, nil)
	if err != nil {
		return err
	}
	return s.client.Login(email, password)
}

func (s *yaService) Disconnect() error {
	return s.client.Logout()
}
