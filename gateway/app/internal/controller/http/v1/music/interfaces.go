package music

import (
	"context"

	"github.com/VrMolodyakov/vgm/gateway/internal/domain/music/model"
)

type AlbumService interface {
	CreateAlbum(ctx context.Context, album model.Album) error
	FindAllAlbums(
		ctx context.Context,
		pagination model.Pagination,
		titleView model.AlbumTitleView,
		releaseView model.AlbumReleasedView,
		sort model.Sort) ([]model.AlbumPreview, error)
	CreatePerson(context.Context, model.Person) error
	FindFullAlbum(ctx context.Context, id string) (model.FullAlbum, error)
	FindLastUpdateDays(ctx context.Context, count uint64) ([]int64, error)
	FindAllPersons(
		ctx context.Context,
		pagination model.Pagination,
		firstNameView model.FirstNameView,
		lastNameView model.LastNameView,
	) ([]model.Person, error)
}
