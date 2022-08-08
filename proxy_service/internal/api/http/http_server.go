package http

import (
	"fmt"
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/app"
)

type HTTPServer struct {
	cmdController cmdController
}

func NewHTTPSrv() *HTTPServer {
	return &HTTPServer{
		cmdController: cmdController{mediator: app.ProxyMediator{}},
	}
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
