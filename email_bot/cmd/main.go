package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Guise322/ozon-exercises/email_bot/config"
	"github.com/Guise322/ozon-exercises/email_bot/internal/api"
)

func main() {
	log.Fatalf("%v", runServer())
}

func runServer() error {
	conf, err := config.ReadConfig()
	if err != nil {
		return err
	}
	address := fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer lis.Close()
	log.Printf("the server listening at %v", lis.Addr())
	return api.RunGRPCSrv(lis)
}
