package mediator

import "reflect"

func (m *CmdMediator) RegHandler(cmd interface{}, handl interface{}) {
	m.handlers[reflect.TypeOf(cmd)] = handl
}
