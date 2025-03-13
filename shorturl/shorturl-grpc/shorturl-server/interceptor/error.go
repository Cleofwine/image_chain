package interceptor

import (
	"context"
	"shorturl-grpc/pkg/xerror"

	"google.golang.org/grpc"
)

func UnaryErrorInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	resp, err = handler(ctx, req)
	switch err.(type) {
	case *xerror.Error:
		err = xerror.New("触发了业务流程限制")
	}
	return resp, err
}

func StreamErrorInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, ss)
	switch err.(type) {
	case *xerror.Error:
		err = xerror.New("触发了业务流程限制")
	}
	return err
}
