package contract

import "context"

type NotifCmd struct {
	Ctx     context.Context
	Id      int64
	From    string
	Message string
}
