package main

import (
	"log"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/email_service/internal/api/server"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/email_service/internal/app/handler"
	"github.com/Guise322/ozon-exercises/email_service/internal/infra"
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
	cl, err := infra.NewNotifClient(path)
	if err != nil {
		return err
	}
	med := mediator.NewMediator()
	med.RegHandler(contract.SubscribtionCmd{}, &handler.SubCmdHandler{NotifClient: cl})
	med.RegHandler(contract.UnreadCountRequest{}, &handler.EmailReqHandler{})
	return server.RunGRPCSrv(path, med)
}

func getPath() string {
	if common.IsDebugging() {
		return debugPath
	}
	return prodPath
}
