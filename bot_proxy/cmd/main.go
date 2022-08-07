package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/common/email_service_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf, err := common.ReadConfig()
	if err != nil {
		log.Fatalf("error of reading the config file: %v", err)
	}
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error of creating a gRPC client: %v", err)
	}
	defer conn.Close()
	client := email_service_pb.NewEmailClient(conn)
	resp, _ := client.SubscribeToInbox(context.Background(), &email_service_pb.SubscribeCom{})
	fmt.Printf("the response: %v", resp.Result)
}
