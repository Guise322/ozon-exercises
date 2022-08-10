package main

import (
	"log"

	"github.com/Guise322/ozon-exercises/common"
	"github.com/Guise322/ozon-exercises/email_service/internal/api/server"
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
	return server.RunGRPCSrv(path)
}

func getPath() string {
	if common.IsDebugging() {
		return debugPath
	}
	return prodPath
}
