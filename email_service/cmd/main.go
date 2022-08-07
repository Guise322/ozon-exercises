package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/email_service/internal/api"
)

const (
	debugPath string = "../config/config.yml"
	prodPath  string = "email_service/config/config.yml"
)

func main() {
	log.Fatal(runServer())
}

func runServer() error {
	var p string
	if common.IsDebugging() {
		p = debugPath
	} else {
		p = prodPath
	}
	conf, err := api.ReadConfig(p)
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
	return api.RunGRPCSrv(conf.Server.TimeoutInMils, lis)
}
