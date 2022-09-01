package imap_service

import "time"

type Message struct {
	From *string
	Date time.Time
}
