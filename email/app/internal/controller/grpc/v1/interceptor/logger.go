package interceptor

import (
	"context"

	"github.com/VrMolodyakov/vgm/email/app/pkg/logging"
	"google.golang.org/grpc"
)

func NewLoggerInterceptor(l logging.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		l.Info(info.FullMethod)
		return handler(ctx, req)
	}
}
