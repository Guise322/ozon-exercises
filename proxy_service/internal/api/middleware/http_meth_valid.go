package middleware

type HTTPMethValid struct {
	method string
}

func NewHTTPValidator(method string) *HTTPMethValid {
	return &HTTPMethValid{method: method}
}
