package unread_cnt_client

import (
	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/conf"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/infra/grpc_conn"
)

type unreadCntClient struct {
	emailServiceClient pb.EmailServiceClient
}

func NewUnreadCntClient(c conf.UnreadCntClientConf) (*unreadCntClient, error) {
	conn, err := grpc_conn.CreatGRPCConn(c.UnreadCntClient.Host, c.UnreadCntClient.Port)
	if err != nil {
		return nil, err
	}
	grpcClient := pb.NewEmailServiceClient(conn)
	return &unreadCntClient{emailServiceClient: grpcClient}, nil
}

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
