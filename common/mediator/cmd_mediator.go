package mediator

type CmdMediator struct {
	handlers map[interface{}]interface{}
}

func NewMediator() *CmdMediator {
	return &CmdMediator{handlers: make(map[interface{}]interface{})}
}
