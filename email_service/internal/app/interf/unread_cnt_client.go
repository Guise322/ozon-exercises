package interf

import (
	"context"

	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

type UnreadCntClient interface {
	GetUnreadEmailCnt(ctx context.Context, req *contract.UnreadCountRequest) (*contract.UnreadReqResult, error)
}
