package http_server

import (
	"net/http"

	"github.com/Guise322/ozon-exercises/proxy_service/internal/api/middleware"
)

func (s *httpServer) UseRoutes() {
	subHandler := http.HandlerFunc(s.subscribeToInbox)
	subValidator := middleware.NewHTTPValidator("POST")
	subHandlerMid := http.Handler(subValidator.ValidateHTTPMethod(subHandler))
	http.Handle("/subscribe", subHandlerMid)
	unreadCntHandler := http.HandlerFunc(s.getUnreadEmailCnt)
	unreadCntValidator := middleware.NewHTTPValidator("GET")
	unreadCntHandlerMid := http.Handler(unreadCntValidator.ValidateHTTPMethod(unreadCntHandler))
	http.Handle("/unreadCount", unreadCntHandlerMid)
}
