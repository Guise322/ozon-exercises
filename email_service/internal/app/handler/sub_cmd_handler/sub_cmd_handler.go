package sub_cmd_handler

import "github.com/Guise322/ozon-exercises/email_service/internal/app/interf"

type subCmdHandler struct {
	cl interf.SubClient
}

func NewSubCmdHandler(cl interf.SubClient) *subCmdHandler {
	return &subCmdHandler{cl: cl}
}
