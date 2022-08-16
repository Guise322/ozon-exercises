package interf

import (
	"context"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type UnreadCntClient interface {
	GetUnreadEmailCnt(ctx context.Context, req *contract.UnreadCntReq) (*contract.UnreadCntResult, error)
}
