package http_server

import (
	"fmt"
	"net/http"

	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/proxy_service/internal"
)

func RunHTTPSrv(confPath string, med mediator.Mediator) error {
	var conf HTTPConf
	err := internal.ReadConfig(&conf, confPath)
	if err != nil {
		return err
	}
	httpServ := newHTTPServer(med)
	address := fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)
	httpServ.UseRoutes()
	err = http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}
