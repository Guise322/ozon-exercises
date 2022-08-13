package contract

import "context"

type UnreadCntReq struct {
	Ctx   context.Context
	Login string
	Pass  string
}
