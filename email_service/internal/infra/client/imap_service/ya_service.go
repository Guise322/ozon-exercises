package imap_service

import (
	"errors"
	"log"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type yaService struct {
	client    *client.Client
	localDone chan struct{}
	mesCh     chan<- *imap.Message
}

func NewYaService(mesCh chan<- *imap.Message) *yaService {
	return &yaService{mesCh: mesCh}
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

func (s *yaService) GetUnreadMessageCount() (int, error) {
	_, err := s.client.Select("INBOX", true)
	if err != nil {
		return 0, err
	}
	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{"\\Seen"}
	uids, err := s.client.Search(criteria)
	if err != nil {
		return 0, err
	}
	return len(uids), nil
}

func (s *yaService) SubsribeToInbox(unsubCh chan<- struct{}) error {
	if s.localDone != nil {
		return errors.New("the client already subscribed")
	}
	s.localDone = make(chan struct{})
	go func() {
		idleTimer := time.NewTicker(20 * time.Minute)
		s.client.Select("INBOX", true)
		updates := make(chan client.Update, 10)
		s.client.Updates = updates
		defer func() {
			s.client.Unselect()
			idleTimer.Stop()
			unsubCh <- struct{}{}
		}()
		idleWait := make(chan struct{}, 1)
		firstResp := false
		var idleDone chan struct{}
		for {
			idleDone = make(chan struct{}, 1)
			go func() {
				s.client.Idle(idleDone, nil)
				idleWait <- struct{}{}
			}()
			select {
			case update := <-updates:
				if mbox, ok := update.(*client.MailboxUpdate); ok {
					close(idleDone)
					<-idleWait
					if !firstResp {
						s.mesCh <- <-s.getMessages(mbox)
						firstResp = true
					} else {
						firstResp = false
					}
				}
			case <-idleTimer.C:
				log.Println("Reset Idle")
				close(idleDone)
			case <-s.localDone:
				log.Println("Not idling anymore")
				close(idleDone)
				<-idleWait
				s.localDone = nil
				s.client.Updates = nil
				return
			}
		}
	}()
	return nil
}

func (s *yaService) getMessages(mbox *client.MailboxUpdate) <-chan *imap.Message {
	msgNum := mbox.Mailbox.Messages
	seqSet := new(imap.SeqSet)
	seqSet.AddNum(msgNum)
	messages := make(chan *imap.Message)
	go func() {
		s.client.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()
	return messages
}

func (s *yaService) UnsubscribeInbox(unsubCh <-chan struct{}) error {
	if s.localDone == nil {
		return errors.New("the client is not subscribed")
	}
	s.localDone <- struct{}{}
	<-unsubCh
	return nil
}
