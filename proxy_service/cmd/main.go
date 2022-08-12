package main

import (
	"log"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/grpc_server"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/http_server"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/infra"
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
	return grpc_server.RunGRPCSrv(path)
}

func runHTTPServer() error {
	path := getConfPath()
	cl, err := infra.NewSubClient(path)
	if err != nil {
		return err
	}
	return http_server.RunHTTPSrv(path, cl)
}
