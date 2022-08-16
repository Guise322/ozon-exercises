package mediator

import "context"

type Mediator interface {
	Handle(ctx context.Context, msg interface{}) (interface{}, error)
}
