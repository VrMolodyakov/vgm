package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/youtube/pkg/errors"
	"github.com/VrMolodyakov/vgm/youtube/pkg/logging"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	tracer = otel.Tracer("music-service")
)

type MusicGrpcClient interface {
	FindRandomTitles(ctx context.Context, count uint64) ([]string, error)
}

type music struct {
	client MusicGrpcClient
	logger logging.Logger
}

func NewMusicService(client MusicGrpcClient, logger logging.Logger) *music {
	return &music{
		client: client,
		logger: logger,
	}
}

func (m *music) FindRandomTitles(ctx context.Context, count uint64) ([]string, error) {
	ctx, span := tracer.Start(ctx, "client.FindRandomTitles")
	defer span.End()

	fullAlbum, err := m.client.FindRandomTitles(ctx, count)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				m.logger.Error("Has Internal Error")
				return nil, errors.NewInternal(err, "grpc client response")
			case codes.Aborted:
				m.logger.Error("gRPC Aborted the call")
				return nil, err
			default:
				m.logger.Error(e.Code(), e.Message())
				return nil, err
			}
		}
		return nil, err
	}
	return fullAlbum, nil
}
