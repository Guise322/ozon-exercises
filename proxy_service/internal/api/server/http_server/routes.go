package http_server

import (
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/middleware/http_mid/meth_valid"
	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/middleware/http_mid/timeout_mid"
)

func (s *httpServer) UseRoutes(timeout int64) {
	subHandler := http.HandlerFunc(s.SubscribeToInbox)
	subValidator := meth_valid.NewHTTPValidator(http.MethodGet)
	timeoutMid := timeout_mid.NewHTTPTimeout(timeout)
	subWithMid := http.Handler(subValidator.ValidateHTTPMethod(
		timeoutMid.SetReqTimeout(subHandler)))
	http.Handle("/subscribe", subWithMid)
	unreadCntHandler := http.HandlerFunc(s.GetUnreadEmailCnt)
	unreadCntValidator := meth_valid.NewHTTPValidator(http.MethodGet)
	unreadWithMid := http.Handler(unreadCntValidator.ValidateHTTPMethod(
		timeoutMid.SetReqTimeout(unreadCntHandler)))
	http.Handle("/unreadCount", unreadWithMid)
}
