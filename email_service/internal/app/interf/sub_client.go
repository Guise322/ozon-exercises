package interf

import "github.com/Guise322/ozon-exercises/email_service/internal/app/contract"

type SubClient interface {
	SubToInbox(cmd *contract.SubscribtionCmd) (*contract.SubCmdResult, error)
}
