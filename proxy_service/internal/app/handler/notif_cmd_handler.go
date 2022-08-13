package handler

import (
	"github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"
	"github.com/golang/protobuf/ptypes/empty"
)

type NotifCmdHandler struct{}

func (NotifCmdHandler) Handle(cmd contract.NotifCmd) (empty.Empty, error) {
	return empty.Empty{}, nil
}
