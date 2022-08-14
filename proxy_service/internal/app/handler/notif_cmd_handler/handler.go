package notif_cmd_handler

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/golang/protobuf/ptypes/empty"
)

func (*notifCmdHandler) Handle(cmd contract.NotifCmd) (empty.Empty, error) {
	return empty.Empty{}, nil
}
