package http_server

import (
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/middleware/http_mid"
)

func (s *httpServer) UseRoutes(timeout int64) {
	subHandler := http.HandlerFunc(s.subscribeToInbox)
	subValidator := http_mid.NewHTTPValidator(http.MethodPost)
	timeoutMid := http_mid.NewHTTPTimeout(timeout)
	subWithMid := http.Handler(subValidator.ValidateHTTPMethod(
		timeoutMid.SetReqTimeout(subHandler)))
	http.Handle("/subscribe", subWithMid)
	unreadCntHandler := http.HandlerFunc(s.getUnreadEmailCnt)
	unreadCntValidator := http_mid.NewHTTPValidator(http.MethodGet)
	unreadWithMid := http.Handler(unreadCntValidator.ValidateHTTPMethod(
		timeoutMid.SetReqTimeout(unreadCntHandler)))
	http.Handle("/unreadCount", unreadWithMid)
}
