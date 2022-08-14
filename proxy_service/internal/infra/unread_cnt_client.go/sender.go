package unread_cnt_client

import (
	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
)

func (c *unreadCntClient) GetUnreadEmailCnt(req *contract.UnreadCntReq) (*contract.UnreadCntResult, error) {
	res, err := c.emailServiceClient.GetUnreadEmailCount(req.Ctx, &pb.UnreadCountRequest{
		EmailLogin: req.Login,
		EmailPass:  req.Pass,
	})
	if err != nil {
		return nil, err
	}
	return &contract.UnreadCntResult{Count: res.MessageCount, Error: res.Error}, nil
}
