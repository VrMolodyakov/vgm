package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/music/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	tracer = otel.Tracer("music-service")
)

type MusicGrpcClient interface {
	CreateAlbum(context.Context, model.Album) error
	FindAll(
		ctx context.Context,
		pagination model.Pagination,
		titleView model.AlbumTitleView,
		releaseView model.AlbumReleasedView,
		sort model.Sort) ([]model.AlbumView, error)
	CreatePerson(context.Context, model.Person) error
	FindFullAlbum(ctx context.Context, id string) (model.FullAlbum, error)
}

type music struct {
	client MusicGrpcClient
}

func NewAlbumService(client MusicGrpcClient) *music {
	return &music{
		client: client,
	}
}

func (a *music) CreateAlbum(ctx context.Context, album model.Album) error {
	ctx, span := tracer.Start(ctx, "client.CreateAlbum")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	err := a.client.CreateAlbum(ctx, album)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				logger.Error("Has Internal Error")
				return errors.NewInternal(err, "grpc client response")
			case codes.Aborted:
				logger.Error("gRPC Aborted the call")
				return err
			default:
				logger.Sugar().Error(e.Code(), e.Message())
				return err
			}
		}
		return err
	}

	return nil
}

func (a *music) CreatePerson(ctx context.Context, person model.Person) error {
	ctx, span := tracer.Start(ctx, "client.CreatePerson")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	err := a.client.CreatePerson(ctx, person)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				logger.Error("Has Internal Error")
				return errors.NewInternal(err, "grpc client response")
			case codes.Aborted:
				logger.Error("gRPC Aborted the call")
				return err
			default:
				logger.Sugar().Error(e.Code(), e.Message())
				return err
			}
		}
		return err
	}

	return nil
}

func (m *music) FindAllAlbums(
	ctx context.Context,
	pagination model.Pagination,
	titleView model.AlbumTitleView,
	releaseView model.AlbumReleasedView,
	sort model.Sort) ([]model.AlbumView, error) {

	ctx, span := tracer.Start(ctx, "client.FindAllAlbums")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	albums, err := m.client.FindAll(ctx, pagination, titleView, releaseView, sort)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				logger.Error("Has Internal Error")
				return nil, errors.NewInternal(err, "grpc client response")
			case codes.Aborted:
				logger.Error("gRPC Aborted the call")
				return nil, err
			default:
				logger.Sugar().Error(e.Code(), e.Message())
				return nil, err
			}
		}
		return nil, err
	}
	return albums, nil
}

func (m *music) FindFullAlbum(ctx context.Context, id string) (model.FullAlbum, error) {
	ctx, span := tracer.Start(ctx, "client.FindFullAlbum")
	defer span.End()

	logger := logging.LoggerFromContext(ctx)
	fullAlbum, err := m.client.FindFullAlbum(ctx, id)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Internal:
				logger.Error("Has Internal Error")
				return model.FullAlbum{}, errors.NewInternal(err, "grpc client response")
			case codes.Aborted:
				logger.Error("gRPC Aborted the call")
				return model.FullAlbum{}, err
			default:
				logger.Sugar().Error(e.Code(), e.Message())
				return model.FullAlbum{}, err
			}
		}
		return model.FullAlbum{}, err
	}
	return fullAlbum, nil
}
