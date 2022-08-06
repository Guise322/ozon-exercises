package contracts

type EmailRequest struct {
	Id int64
}

type EmailReqResult struct {
	Id   int64
	From string
	To   string
	Text string
}

type EmailReqHandler interface {
	Handle() *EmailReqResult
}
