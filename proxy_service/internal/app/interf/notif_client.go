package interf

import "github.com/Guise322/ozon-exercises/proxy_service/internal/app/contract"

type SubClient interface {
	SubToInbox(cmd *contract.ProxySubCmd) error
}
