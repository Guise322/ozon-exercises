package ya_imap

import "github.com/emersion/go-imap"

func (s *yaService) GetUnreadMessageCount() (int64, error) {
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
	return int64(len(uids)), nil
}
