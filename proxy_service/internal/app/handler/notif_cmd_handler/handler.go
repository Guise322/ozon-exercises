package notif_cmd_handler

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/golang/protobuf/ptypes/empty"
)

func (h *notifCmdHandler) Handle(cmd *contract.NotifCmd) (empty.Empty, error) {
	_, err := h.cl.Notify(cmd)
	if err != nil {
		return empty.Empty{}, err
	}
	return empty.Empty{}, nil
}
