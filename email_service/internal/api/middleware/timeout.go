package middleware

type timeoutInterceptor struct {
	timeout int64
}

func NewTimeoutInterc(timeout int64) *timeoutInterceptor {
	return &timeoutInterceptor{timeout: timeout}
}
