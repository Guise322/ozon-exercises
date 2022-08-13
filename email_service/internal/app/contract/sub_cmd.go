package contract

import "context"

type SubscribtionCmd struct {
	Ctx        context.Context
	EmailLogin string
	EmailPass  string
}
