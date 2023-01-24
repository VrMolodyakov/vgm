package model

import (
	"fmt"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/types"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

const (
	nameFilter    = "name"
	publishFilter = "published_at"
	personFilter  = "person"
)

func AlbumFilterFields() map[string]string {
	return map[string]string{
		nameFilter:    filter.DataTypeStr,
		publishFilter: filter.DataTypeDate,
		personFilter:  filter.DataTypeStr,
	}
}

func AlbumFilter(req *albumPb.FindAllAlbumsRequest) filter.Filterable {
	options := filter.NewOptions(req.GetPagination().GetLimit(), req.GetPagination().GetOffset(), AlbumFilterFields())

	if req == nil {
		return options
	}

	name := req.GetName()
	if name != nil {
		operator := types.StringOperatorFromPB(req.GetName().GetOp())
		addFilterField(nameFilter, name.GetVal(), operator, options)
	}

	published := req.GetPublishedAt()
	if published != nil {
		operator := types.IntOperatorFromPB(req.GetPublishedAt().GetOp())
		addFilterField(publishFilter, published.GetVal(), operator, options)
	}

	person := req.GetPerson()
	if person != nil {
		operator := types.StringOperatorFromPB(req.GetPerson().GetOp())
		addFilterField(personFilter, person.GetVal(), operator, options)
	}

	return nil
}

func addFilterField(
	name, value string,
	operator string,
	options filter.Filterable,
) {
	err := options.AddField(name, operator, value)
	if err != nil {
		logging.GetLogger().With(
			err,
			fmt.Errorf("failed to add filter field. name=%s, operator=%s, value=%s",
				name,
				operator,
				value),
		)
	}
}
