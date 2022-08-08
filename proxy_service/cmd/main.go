package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Guise322/ozon-exercises/common"
	grpcSrv "github.com/Guise322/ozon-exercises/proxy_service/internal/api/grpc"
	httpSrv "github.com/Guise322/ozon-exercises/proxy_service/internal/api/http"
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

func runGRPCServer() error {
	var p string
	if common.IsDebugging() {
		p = debugPath
	} else {
		p = prodPath
	}
	conf, err := grpcSrv.ReadConfig(p)
	if err != nil {
		return err
	}
	address := fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer lis.Close()
	return grpcSrv.RunGRPCSrv(lis)
}

func runHTTPServer() error {
	var p string
	if common.IsDebugging() {
		p = debugPath
	} else {
		p = prodPath
	}
	srv := httpSrv.NewHTTPSrv()
	conf, err := httpSrv.ReadConfig(p)
	if err != nil {
		return err
	}
	return srv.RunHTTPSrv(conf)
}
