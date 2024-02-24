package middleware

import (
	"context"
	"google.golang.org/grpc"
)

func (e *ErrsParser) CustomErrsGrpcInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		customErr, isCustom := e.Parse(err.Error())
		if isCustom {
			return customErr
		}
	}

	return err
}
