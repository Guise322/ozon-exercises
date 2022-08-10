package http_server

import (
	"fmt"
	"net/http"
)

func RunHTTPSrv(confPath string) error {
	httpServ := newHTTPServer()
	conf, err := readConfig(confPath)
	if err != nil {
		return err
	}
	address := fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)
	httpServ.UseRoutes()
	err = http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}
