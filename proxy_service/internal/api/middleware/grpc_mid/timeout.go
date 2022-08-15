package grpc_mid

type timeoutInterceptor struct {
	timeout int64
}

func NewTimeoutInterc(timeout int64) *timeoutInterceptor {
	return &timeoutInterceptor{timeout: timeout}
}
