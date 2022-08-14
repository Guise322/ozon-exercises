package interf

import "github.com/Guise322/ozon-exercises/email_service/internal/app/contract"

type NotifClient interface {
	SendNotif(cmd *contract.NotifCmd) error
}
