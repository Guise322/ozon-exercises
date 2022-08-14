package interf

import "github.com/Guise322/ozon-exercises/email_service/internal/app/contract"

type UnreadCntClient interface {
	GetUnreadEmailCnt(req *contract.UnreadCountRequest) (*contract.UnreadReqResult, error)
}
