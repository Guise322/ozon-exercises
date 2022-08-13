package mediator

type Mediator interface {
	Handle(msg interface{}) (interface{}, error)
}
