package model

import (
	"fmt"

	albumPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/types"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
	"github.com/VrMolodyakov/vgm/music/app/pkg/sort"
)

const (
	nameFilterType    = filter.DataTypeStr
	publishFilterType = filter.DataTypeDate
	personFilterType  = filter.DataTypeStr

	nameFilter    = "name"
	publishFilter = "published_at"
	personFilter  = "person"
)

func AlbumSort(req *albumPb.FindAllAlbumsRequest) sort.Sortable {
	field := req.GetSort().GetField()
	return sort.NewOptions(field)
}

func AlbumFilter(req *albumPb.FindAllAlbumsRequest) filter.Filterable {
	options := filter.NewOptions(req.GetPagination().GetLimit(), req.GetPagination().GetOffset())

	if req == nil {
		return options
	}

	name := req.GetName()
	if name != nil {
		operator := types.StringOperatorFromPB(req.GetName().GetOp())
		addFilterField(nameFilter, name.GetVal(), operator, nameFilterType, options)
	}

	published := req.GetPublishedAt()
	if published != nil {
		operator := types.IntOperatorFromPB(req.GetPublishedAt().GetOp())
		addFilterField(publishFilter, published.GetVal(), operator, publishFilterType, options)
	}

	person := req.GetPerson()
	if person != nil {
		operator := types.StringOperatorFromPB(req.GetPerson().GetOp())
		addFilterField(personFilter, person.GetVal(), operator, personFilterType, options)
	}

	return options
}

func addFilterField(
	name string,
	value string,
	operator string,
	fieldType string,
	options filter.Filterable,
) {
	err := options.AddField(name, operator, value, fieldType)
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
