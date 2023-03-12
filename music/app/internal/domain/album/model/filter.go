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
	releaseFilterType = filter.DataTypeDate
	personFilterType  = filter.DataTypeStr

	titleFilter   = "title"
	releaseFilter = "released_at"
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

	name := req.GetTitle()
	if name != nil {
		operator := types.StringOperatorFromPB(req.GetTitle().GetOp())
		addFilterField(titleFilter, name.GetVal(), operator, nameFilterType, options)
	}

	released := req.GetReleasedAt()
	if released != nil {
		operator := types.IntOperatorFromPB(req.GetReleasedAt().GetOp())
		addFilterField(releaseFilter, released.GetVal(), operator, releaseFilterType, options)
	}
	fmt.Println("------------Album-FIlter-Model-----------")
	fmt.Println(options)
	fmt.Println("------------Album-FIlter-Model-----------")
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
		logging.GetLogger().Infow(
			err.Error(),
			fmt.Errorf("failed to add filter field. name=%s, operator=%s, value=%s",
				name,
				operator,
				value),
		)
	}
}
