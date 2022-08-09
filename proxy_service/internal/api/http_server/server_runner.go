package http_server

import (
	"fmt"
	"net/http"
)

func (srv *HTTPServer) RunHTTPSrv(conf *httpConf) error {
	address := fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)
	srv.UseRoutes()
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}
