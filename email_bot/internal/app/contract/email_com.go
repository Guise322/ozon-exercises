package contract

type EmailCommand struct {
}

type EmailComResult struct {
	Data string
}

type EmailComHandler interface {
	Handle() (*EmailComResult, error)
}
