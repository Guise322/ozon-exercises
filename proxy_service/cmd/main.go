package main

import (
	"log"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/grpc_server"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/http_server"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/handler/notif_cmd_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/handler/sub_cmd_handler"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/handler/unread_cnt_handler"
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
	// conn, err := grpc.Dial(
	// 	fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalf("error of creating a gRPC client: %v", err)
	// }
	// defer conn.Close()
	// client := email_service_pb.NewEmailClient(conn)
	// resp, _ := client.SubscribeToInbox(context.Background(), &email_service_pb.SubscribeCom{})
	// fmt.Printf("the response: %v", resp.Result)
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
	notifCl, err := notif_client.NewNotifClient(path)
	if err != nil {
		return err
	}
	grpcMed := mediator.NewMediator()
	grpcMed.RegHandler(&contract.NotifCmd{}, notif_cmd_handler.NewNotifCmdHandler(notifCl))
	return grpc_server.RunGRPCSrv(path, grpcMed)
}

func runHTTPServer() error {
	path := getConfPath()
	subCl, err := sub_client.NewSubClient(path)
	if err != nil {
		return err
	}
	unreadCl, err := unread_cnt_client.NewUnreadCntClient(path)
	if err != nil {
		return err
	}
	httpMed := mediator.NewMediator()
	httpMed.RegHandler(&contract.ProxySubCmd{}, sub_cmd_handler.NewSubCmdHandler(subCl))
	httpMed.RegHandler(&contract.UnreadCntReq{}, unread_cnt_handler.NewUnreadCntHandler(unreadCl))
	return http_server.RunHTTPSrv(path, httpMed)
}
