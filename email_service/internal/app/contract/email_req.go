package contract

type EmailRequest struct {
	Id int64
}

type EmailReqResult struct {
	Id   int64
	From string
	To   string
	Text string
}
