package main

import (
	"log"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/server/grpc_server"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/server/http_server"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/cmd_handler/notif_cmd_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/cmd_handler/sub_cmd_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/cmd_handler/unread_cnt_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/conf"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/infra/client/notif"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/infra/client/sub"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/infra/client/unread_cnt.go"
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
	notifCl := notif.NewNotifClient(notifConf)
	grpcMed := mediator.NewMediator()
	grpcMed.RegHandler(&contract.NotifCmd{}, notif_cmd_handler.NewNotifCmdHandler(notifCl))
	var grpcConf conf.GRPCConf
	err = conf.ReadConfig(&grpcConf, path)
	if err != nil {
		return err
	}
	return grpc_server.NewGRPCServer(grpcMed).RunGRPCSrv(grpcConf)
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
	subCl, err := sub.NewSubClient(subConf)
	if err != nil {
		return err
	}
	unreadCl, err := unread_cnt.NewUnreadCntClient(unreadConf)
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
	return http_server.NewHTTPServer(httpMed).RunHTTPSrv(httpConf)
}
