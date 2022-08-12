package handler

import (
	"context"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

type UnreadCntHandler struct{}

func (UnreadCntHandler) Handle(ctx context.Context, req contract.UnreadCntReq) (contract.UnreadCntResult, error) {
	runes := []rune(req.Pass)
	r := runes[0]
	return contract.UnreadCntResult{Count: int64(r)}, nil
}
