package contracts

type EmailRequest struct {
	Id int64
}

type EmailResponse struct {
	Id   int64
	From string
	To   string
	Text string
}

type EmailReqHandler interface {
	Handle(EmailRequest) EmailResponse
}
