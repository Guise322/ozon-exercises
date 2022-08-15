package main

import (
	"log"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/email_service/internal/api/server"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/handler/sub_cmd_handler"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/handler/unread_cnt_handler"
	"github.com/Guise322/ozon-exercises/email_service/internal/infra/sub_client"
	"github.com/Guise322/ozon-exercises/email_service/internal/infra/unread_cnt_client"
)

const (
	debugPath string = "../config/config.yml"
	prodPath  string = "email_service/config/config.yml"
)

func main() {
	log.Fatal(runServer())
}

func runServer() error {
	path := getPath()
	med := mediator.NewMediator()
	subClient, err := sub_client.NewSubClient()
	if err != nil {
		return err
	}
	unreadCntClient, err := unread_cnt_client.NewUnreadCntClient()
	if err != nil {
		return err
	}
	med.RegHandler(&contract.SubscribtionCmd{}, sub_cmd_handler.NewSubCmdHandler(subClient))
	med.RegHandler(&contract.UnreadCountRequest{}, unread_cnt_handler.NewUnreadCntHandler(unreadCntClient))
	return server.RunGRPCSrv(path, med)
}

func getPath() string {
	if common.IsDebugging() {
		return debugPath
	}
	return prodPath
}
