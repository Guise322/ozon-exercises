package main

import (
	"log"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/grpc_server"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/http_server"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/cmd_handler/notif_cmd_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/cmd_handler/sub_cmd_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/cmd_handler/unread_cnt_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/conf"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/infra/notif_client"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/infra/sub_client"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/infra/unread_cnt_client.go"
)

const (
	debugPath string = "../config/config.yml"
	prodPath  string = "proxy_service/config/config.yml"
)

func main() {
	go func() { log.Print(runGRPCServer()) }()
	log.Fatal(runHTTPServer())
}

func getConfPath() string {
	if common.IsDebugging() {
		return debugPath
	} else {
		return prodPath
	}
}

func runGRPCServer() error {
	path := getConfPath()
	var notifConf conf.NotifClientConf
	err := conf.ReadConfig(&notifConf, path)
	if err != nil {
		return err
	}
	notifCl, err := notif_client.NewNotifClient(notifConf)
	if err != nil {
		return err
	}
	grpcMed := mediator.NewMediator()
	grpcMed.RegHandler(&contract.NotifCmd{}, notif_cmd_handler.NewNotifCmdHandler(notifCl))
	var grpcConf conf.GRPCConf
	err = conf.ReadConfig(&grpcConf, path)
	if err != nil {
		return err
	}
	return grpc_server.RunGRPCSrv(grpcConf, grpcMed)
}

func runHTTPServer() error {
	path := getConfPath()
	var subConf conf.SubClientConf
	err := conf.ReadConfig(&subConf, path)
	if err != nil {
		return err
	}
	var unreadConf conf.UnreadCntClientConf
	err = conf.ReadConfig(&unreadConf, path)
	if err != nil {
		return err
	}
	subCl, err := sub_client.NewSubClient(subConf)
	if err != nil {
		return err
	}
	unreadCl, err := unread_cnt_client.NewUnreadCntClient(unreadConf)
	if err != nil {
		return err
	}
	httpMed := mediator.NewMediator()
	httpMed.RegHandler(&contract.ProxySubCmd{}, sub_cmd_handler.NewSubCmdHandler(subCl))
	httpMed.RegHandler(&contract.UnreadCntReq{}, unread_cnt_handler.NewUnreadCntHandler(unreadCl))
	var httpConf conf.HTTPConf
	err = conf.ReadConfig(&httpConf, path)
	if err != nil {
		return err
	}
	return http_server.RunHTTPSrv(httpConf, httpMed)
}
