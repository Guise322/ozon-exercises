package http_server

import (
	"fmt"
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/conf"
)

func (s *httpServer) RunHTTPSrv(c conf.HTTPConf) error {
	address := fmt.Sprintf("%v:%v", c.Server.Host, c.Server.Port)
	s.UseRoutes(c.Server.Timeout)
	return http.ListenAndServe(address, nil)
}
