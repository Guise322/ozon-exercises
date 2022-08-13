package contract

import "context"

type ProxySubCmd struct {
	Ctx   context.Context
	Login string
	Pass  string
}
