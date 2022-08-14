package interf

import "github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"

type UnreadCntClient interface {
	GetUnreadEmailCnt(req *contract.UnreadCntReq) (*contract.UnreadCntResult, error)
}
