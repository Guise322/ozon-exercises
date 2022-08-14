package notif_cmd_handler

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
)

type notifCmdHandler struct {
	cl interf.NotifClient
}

func NewNotifCmdHandler(cl interf.NotifClient) *notifCmdHandler {
	return &notifCmdHandler{cl: cl}
}
