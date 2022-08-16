package http_server

import (
	"fmt"
	"net/http"

	"github.com/Guise322/ozon-exercises/common/mediator"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/conf"
)

func RunHTTPSrv(c conf.HTTPConf, med mediator.Mediator) error {
	httpServ := newHTTPServer(med)
	address := fmt.Sprintf("%v:%v", c.Server.Host, c.Server.Port)
	httpServ.UseRoutes(c.Server.Timeout)
	return http.ListenAndServe(address, nil)
}
