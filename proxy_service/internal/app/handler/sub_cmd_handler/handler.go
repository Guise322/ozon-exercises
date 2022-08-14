package sub_cmd_handler

import "github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"

func (h *subCmdHandler) Handle(cmd *contract.ProxySubCmd) (interface{}, error) {
	return nil, h.subClient.SubToInbox(cmd)
}
