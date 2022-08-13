package contract

import "context"

type UnreadCountRequest struct {
	Ctx        context.Context
	EmailLogin string
	EmailPass  string
}
