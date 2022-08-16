package mediator

import (
	"context"
	"errors"
	"reflect"
)

type CmdMediator struct {
	handlers map[interface{}]interface{}
}

func NewMediator() *CmdMediator {
	return &CmdMediator{handlers: make(map[interface{}]interface{})}
}

func (m *CmdMediator) RegHandler(msg interface{}, handl interface{}) {
	m.handlers[reflect.TypeOf(msg)] = handl
}

func (m *CmdMediator) Handle(ctx context.Context, msg interface{}) (interface{}, error) {
	handl, ok := m.handlers[reflect.TypeOf(msg)]
	if !ok {
		return nil, errors.New("have no handler for the command")
	}
	v := reflect.ValueOf(handl)
	meth := v.MethodByName("Handle")
	if !meth.IsValid() {
		return nil, errors.New("have no method with the name \"Handle\"")
	}
	vals := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(msg)}
	results := meth.Call(vals)
	if !results[1].IsNil() {
		return nil, results[1].Interface().(error)
	}
	return results[0].Interface(), nil
}
