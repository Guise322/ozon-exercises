package sub_cmd_handler

import (
	"github.com/Guise322/ozon-exercises/email_service/internal/app/contract"
)

func (h *subCmdHandler) Handle(cmd *contract.SubscribtionCmd) (*contract.SubCmdResult, error) {
	// ticker := time.NewTicker(5 * time.Millisecond)
	// defer ticker.Stop()
	// select {
	// case <-ticker.C:
	// 	h.cl.SubToInbox(&contract.SubscribtionCmd{
	// 		Ctx:        cmd.Ctx,
	// 		EmailLogin: cmd.EmailLogin,
	// 		EmailPass:  cmd.EmailPass,
	// 	})
	// case <-cmd.Ctx.Done():
	// 	return nil, errors.New("timeout")
	// }
	// return &contract.SubCmdResult{Error: "Hello there!"}, errors.New("Ho-ho")

	return h.cl.SubToInbox(&contract.SubscribtionCmd{
		Ctx:        cmd.Ctx,
		EmailLogin: cmd.EmailLogin,
		EmailPass:  cmd.EmailPass,
	})
}
