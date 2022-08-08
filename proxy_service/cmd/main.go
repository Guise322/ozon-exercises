package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/grpc_server"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/http_server"
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
	conf, err := grpc_server.ReadConfig(path)
	if err != nil {
		return err
	}
	address := fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer lis.Close()
	return grpc_server.RunGRPCSrv(lis)
}

func runHTTPServer() error {
	path := getConfPath()
	srv := http_server.NewHTTPSrv()
	conf, err := http_server.ReadConfig(path)
	if err != nil {
		return err
	}
	return srv.RunHTTPSrv(conf)
}
