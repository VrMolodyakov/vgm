package types

import (
	filterPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/filter/v1"
	"github.com/VrMolodyakov/vgm/music/app/pkg/filter"
)

func IntOperatorFromPB(e filterPb.IntFieldFilter_Operator) string {
	switch e {
	case filterPb.IntFieldFilter_OPERATOR_EQ:
		return filter.OperatorEq
	case filterPb.IntFieldFilter_OPERATOR_NEQ:
		return filter.OperatorNotEq
	case filterPb.IntFieldFilter_OPERATOR_LT:
		return filter.OperatorLowerThan
	case filterPb.IntFieldFilter_OPERATOR_LTE:
		return filter.OperatorLowerThanEq
	case filterPb.IntFieldFilter_OPERATOR_GT:
		return filter.OperatorGreaterThan
	case filterPb.IntFieldFilter_OPERATOR_GTE:
		return filter.OperatorGreaterThanEq
	default:
		return ""
	}
}

func StringOperatorFromPB(e filterPb.StringFieldFilter_Operator) string {
	switch e {
	case filterPb.StringFieldFilter_OPERATOR_EQ:
		return filter.OperatorEq
	case filterPb.StringFieldFilter_OPERATOR_NEQ:
		return filter.OperatorNotEq
	case filterPb.StringFieldFilter_OPERATOR_LIKE:
		return filter.OperatorLike
	default:
		return ""
	}
}
