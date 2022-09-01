package sub_client

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb/common/proto"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_service/internal/conf"
	"github.com/Guise322/ozon-exercises/email_service/internal/infra"
	"github.com/Guise322/ozon-exercises/email_service/internal/infra/client/grpc/grpc_conn"
	"github.com/Guise322/ozon-exercises/email_service/internal/infra/imap_service"
	"github.com/emersion/go-imap"
)

type subClient struct {
	subService          imap_service.Subscriber
	mesCh               chan *imap.Message
	unsubCh             chan struct{}
	newEmailNotifClient pb.NewEmailNotifClient
	imapAddr            string
}

func NewSubClient(
	subService imap_service.Subscriber,
	c conf.NotifClientConf,
	imapAddr string,
) (*subClient, error) {
	conn, err := grpc_conn.CreatGRPCConn(c.NotifClient.Host, c.NotifClient.Port)
	if err != nil {
		return nil, err
	}
	return &subClient{
		subService:          subService,
		mesCh:               make(chan *imap.Message),
		unsubCh:             make(chan struct{}),
		newEmailNotifClient: pb.NewNewEmailNotifClient(conn),
		imapAddr:            imapAddr,
	}, nil
}

func (c *subClient) SubToInbox(ctx context.Context, cmd *contract.SubscribtionCmd) error {
	if !infra.YandexCheck(cmd.EmailLogin) {
		return errors.New("email_service: the service so far only works with Yandex mail")
	}
	go func() {
		mes := <-c.mesCh
		fromName := mes.Envelope.From[0].PersonalName
		fromAddr := mes.Envelope.From[0].Address()
		grpcMes := &pb.NewEmailCmd{From: fmt.Sprintf("%v %v", fromName, fromAddr)}
		c.newEmailNotifClient.Notify(ctx, grpcMes)
	}()
	if err := c.subService.Connect(c.imapAddr, cmd.EmailLogin, cmd.EmailPass); err != nil {
		return err
	}
	return c.subService.SubToInbox(c.mesCh)
}

func (c *subClient) UnsubFromInbox(ctx context.Context, cmd *contract.SubscribtionCmd) error {
	if !infra.YandexCheck(cmd.EmailLogin) {
		return errors.New("email_service: the service so far only works with Yandex mail")
	}
	if err := c.subService.UnsubFromInbox(); err != nil {
		return err
	}
	return c.subService.Disconnect()
}
