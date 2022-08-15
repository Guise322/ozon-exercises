package http_mid

type HTTPTimeout struct {
	value int64
}

func NewHTTPTimeout(timeout int64) *HTTPTimeout {
	return &HTTPTimeout{value: timeout}
}
