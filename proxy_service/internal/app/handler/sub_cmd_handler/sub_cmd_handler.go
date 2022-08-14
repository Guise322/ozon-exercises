package sub_cmd_handler

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/interf"
)

type subCmdHandler struct {
	subClient interf.SubClient
}

func NewSubCmdHandler(subClient interf.SubClient) *subCmdHandler {
	return &subCmdHandler{subClient: subClient}
}
