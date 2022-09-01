package ya_imap

import (
	"errors"
	"log"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func (s *yaService) SubToInbox(mesCh chan<- *imap.Message) error {
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
			s.unsubCh <- struct{}{}
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
						mesCh <- <-s.getMessages(mbox)
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

func (s *yaService) UnsubFromInbox() error {
	if s.localDone == nil {
		return errors.New("the client is not subscribed")
	}
	s.localDone <- struct{}{}
	<-s.unsubCh
	return nil
}
