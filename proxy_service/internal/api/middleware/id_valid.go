package middleware

import (
	"context"
	"errors"

	pb "github.com/Guise322/ozon-exercises/common/email_service_pb"
	"google.golang.org/grpc"
)

type IdValidInterceptor struct{}

func (v IdValidInterceptor) ValidateId(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	emailCom, ok := req.(*pb.NewEmailCom)
	if !ok {
		return handler(ctx, req)
	}
	if emailCom.Id <= 0 {
		return nil, errors.New("id should be greater than 0")
	}
	return handler(ctx, req)
}
