package interceptor

import (
	"context"
	"errors"
	"shorturl-grpc/pkg/config"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryAuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if info.FullMethod != "/grpc.health.v1.Health/Check" {
		err = oauth2Valid(ctx)
		if err != nil {
			return nil, err
		}
	}
	return handler(ctx, req)
}

func StreamAuthInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := oauth2Valid(ss.Context())
	if err != nil {
		return err
	}
	return handler(srv, ss)
}

func oauth2Valid(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("元数据获取失败，身份认证失败")
	}
	authorization := md["authorization"]
	if len(authorization) < 1 {
		return errors.New("元数据获取失败，身份认证失败")
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	cnf := config.GetConfig()
	if cnf.Server.AccessToken != token {
		return errors.New("鉴权失败")
	}
	return nil
}
