package http

import (
	"fmt"
	"net/http"
)

type HTTPServer struct {
	cmdController cmdController
}

func NewHTTPSrv() *HTTPServer {
	return &HTTPServer{cmdController: cmdController{}}
}

func (srv *HTTPServer) RunHTTPSrv(conf *httpConf) error {
	address := fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port)
	srv.UseRoutes()
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}
