package model

import (
	filterPb "github.com/VrMolodyakov/vgm/music/app/gen/go/proto/filter/v1"
)

type Pagination struct {
	Limit  int
	Offset int
}

type Sort struct {
	Field string
}

type AlbumTitleView struct {
	Value    string
	Operator string
}

type FirstNameView struct {
	Value    string
	Operator string
}

type LastNameView struct {
	Value    string
	Operator string
}

type AlbumReleasedView struct {
	Value    string
	Operator string
}

func (p *Pagination) PbFromModel() *filterPb.Pagination {
	if p.Limit == 0 && p.Offset == 0 {
		return nil
	}
	pb := filterPb.Pagination{}
	if p.Limit != 0 {
		pb.Limit = uint64(p.Limit)
	}
	if p.Offset != 0 {
		pb.Offset = uint64(p.Offset)
	}
	return &pb
}

func (titleView *AlbumTitleView) PbFromModel() *filterPb.StringFieldFilter {
	if titleView.Value == "" {
		return nil
	}
	pb := filterPb.StringFieldFilter{Val: titleView.Value}
	switch op := titleView.Operator; op {
	case "like":
		pb.Op = filterPb.StringFieldFilter_OPERATOR_LIKE
	case "neq":
		pb.Op = filterPb.StringFieldFilter_OPERATOR_NEQ
	case "eq":
		pb.Op = filterPb.StringFieldFilter_OPERATOR_EQ
	default:
		pb.Op = filterPb.StringFieldFilter_OPERATOR_LIKE
	}
	return &pb
}

func (nameView *FirstNameView) PbFromModel() *filterPb.StringFieldFilter {
	if nameView.Value == "" {
		return nil
	}
	pb := filterPb.StringFieldFilter{Val: nameView.Value}
	switch op := nameView.Operator; op {
	case "like":
		pb.Op = filterPb.StringFieldFilter_OPERATOR_LIKE
	case "neq":
		pb.Op = filterPb.StringFieldFilter_OPERATOR_NEQ
	case "eq":
		pb.Op = filterPb.StringFieldFilter_OPERATOR_EQ
	default:
		pb.Op = filterPb.StringFieldFilter_OPERATOR_LIKE
	}
	return &pb
}

func (nameView *LastNameView) PbFromModel() *filterPb.StringFieldFilter {
	if nameView.Value == "" {
		return nil
	}
	pb := filterPb.StringFieldFilter{Val: nameView.Value}
	switch op := nameView.Operator; op {
	case "like":
		pb.Op = filterPb.StringFieldFilter_OPERATOR_LIKE
	case "neq":
		pb.Op = filterPb.StringFieldFilter_OPERATOR_NEQ
	case "eq":
		pb.Op = filterPb.StringFieldFilter_OPERATOR_EQ
	default:
		pb.Op = filterPb.StringFieldFilter_OPERATOR_LIKE
	}
	return &pb
}

func (releaseView *AlbumReleasedView) PbFromModel() *filterPb.IntFieldFilter {
	if releaseView.Value == "" {
		return nil
	}
	pb := filterPb.IntFieldFilter{Val: releaseView.Value}
	switch op := releaseView.Operator; op {
	case "eq":
		pb.Op = filterPb.IntFieldFilter_OPERATOR_EQ
	case "gt":
		pb.Op = filterPb.IntFieldFilter_OPERATOR_GT
	case "gte":
		pb.Op = filterPb.IntFieldFilter_OPERATOR_GTE
	case "lt":
		pb.Op = filterPb.IntFieldFilter_OPERATOR_LT
	case "lte":
		pb.Op = filterPb.IntFieldFilter_OPERATOR_LTE
	case "neq":
		pb.Op = filterPb.IntFieldFilter_OPERATOR_NEQ
	default:
		pb.Op = filterPb.IntFieldFilter_OPERATOR_EQ
	}
	return &pb
}

func (s *Sort) PBFromModel() *filterPb.Sort {
	if s.Field == "" {
		return nil
	}
	return &filterPb.Sort{
		Field: s.Field,
	}
}
