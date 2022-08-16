package interf

import "github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"

type NotifClient interface {
	Notify(cmd *contract.NotifCmd) error
}
