package contracts

type EmailCommand struct {
	Id int64
}

type EmailComHandler interface {
	Handle(EmailCommand) error
}
