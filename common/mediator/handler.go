package mediator

import (
	"errors"
	"reflect"
)

func (m *CmdMediator) Handle(cmd interface{}) (interface{}, error) {
	handl, ok := m.handlers[reflect.TypeOf(cmd)]
	if !ok {
		return nil, errors.New("have no handler for the command")
	}
	v := reflect.ValueOf(handl)
	meth := v.MethodByName("Handle")
	if !meth.IsValid() {
		return nil, errors.New("have no method with the name \"Handle\"")
	}
	vals := []reflect.Value{reflect.ValueOf(cmd)}
	results := meth.Call(vals)
	var err error
	if !results[1].IsNil() {
		err = results[1].Interface().(error)
	}
	return results[0].Interface(), err
}
