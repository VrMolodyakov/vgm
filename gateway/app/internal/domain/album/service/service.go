package service

import (
	"context"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/album/model"
)

type AlbumGrpcClient interface {
	Create(context.Context, model.Album) error
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
	return a.client.Create(ctx, album)
}
