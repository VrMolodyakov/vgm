package model

import (
	"fmt"

	personPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/music_service/album/v1"
	"github.com/VrMolodyakov/vgm/music/app/internal/controller/grpc/types"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
	"github.com/VrMolodyakov/vgm/music/app/pkg/logging"
)

const (
	nameFilterType  = filter.DataTypeStr
	firstNameFilter = "first_name"
	lastNameFilter  = "last_name"
)

func PersonFilter(req *personPb.FindAllPersonsRequest) filter.Filterable {
	options := filter.NewOptions(req.GetPagination().GetLimit(), req.GetPagination().GetOffset())

	if req == nil {
		return options
	}

	firstName := req.GetFirstName()
	if firstName != nil {
		operator := types.StringOperatorFromPB(req.GetFirstName().GetOp())
		addFilterField(firstNameFilter, firstName.GetVal(), operator, nameFilterType, options)
	}

	lastName := req.GetFirstName()
	if lastName != nil {
		operator := types.StringOperatorFromPB(req.GetLastName().GetOp())
		addFilterField(firstNameFilter, lastName.GetVal(), operator, nameFilterType, options)
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
