package model

import (
	"fmt"

	personPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/person/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/types"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

const (
	fullNameFilterType = filter.DataTypeStr
	fullNameFilter     = "title"
)

func PersonFilter(req *personPb.FindAllPersonsRequest) filter.Filterable {
	options := filter.NewOptions(req.GetPagination().GetLimit(), req.GetPagination().GetOffset())

	if req == nil {
		return options
	}

	fullName := req.GetFullName()
	if fullName != nil {
		operator := types.StringOperatorFromPB(req.GetFullName().GetOp())
		addFilterField(fullNameFilter, fullName.GetVal(), operator, fullNameFilterType, options)
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
		logging.GetLogger().Infow(
			err.Error(),
			fmt.Errorf("failed to add filter field. name=%s, operator=%s, value=%s",
				name,
				operator,
				value),
		)
	}
}
