package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/album/model"
	"github.com/VrMolodyakov/vgm/gateway/pkg/errors"
	"github.com/VrMolodyakov/vgm/gateway/pkg/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AlbumGrpcClient interface {
	Create(context.Context, model.Album) error
	FindAll(
		ctx context.Context,
		pagination model.Pagination,
		titleView model.AlbumTitleView,
		releaseView model.AlbumReleasedView,
		sort model.Sort,
	) error
}

type album struct {
	client AlbumGrpcClient
}

func NewAlbumService(client AlbumGrpcClient) *album {
	return &album{
		client: client,
	}
}

func (a *album) CreateAlbum(ctx context.Context, album model.Album) error {
	logger := logging.LoggerFromContext(ctx)
	err := a.client.Create(ctx, album)
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

func (a *album) FindAllAlbums(
	ctx context.Context,
	pagination model.Pagination,
	titleView model.AlbumTitleView,
	releaseView model.AlbumReleasedView,
	sort model.Sort,
) error {

	// logger := logging.LoggerFromContext(ctx)
	a.client.FindAll(ctx, pagination, titleView, releaseView, sort)

	return nil
}
