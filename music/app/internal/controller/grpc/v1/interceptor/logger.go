package interceptor

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewLoggerInterceptor(l *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		l.Info(info.FullMethod)
		return handler(ctx, req)
	}
}
