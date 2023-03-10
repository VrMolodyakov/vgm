package filter

import (
	"fmt"
)

const (
	DataTypeStr   = "string"
	DataTypeInt   = "int"
	DataTypeBool  = "bool"
	DataTypeDate  = "date"
	DataTypeArray = "array"
	DataTypeNull  = "empty"

	OperatorEq            = "eq"
	OperatorNotEq         = "neq"
	OperatorLowerThan     = "lt"
	OperatorLowerThanEq   = "le"
	OperatorGreaterThan   = "gt"
	OperatorGreaterThanEq = "ge"
	OperatorIn            = "in"
	OperatorLike          = "like"
)

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}

type Options struct {
	filter string
	limit  uint64
	offset uint64
	fields []Field
}

type Filterable interface {
	Limit() uint64
	Offset() uint64
	Fields() []Field
	AddField(name string, operator string, value string, fieldType string) error
}

func NewOptions(limit, offset uint64) *Options {
	return &Options{
		limit:  limit,
		offset: offset,
	}
}

func (o *Options) Limit() uint64 {
	return o.limit
}

func (o *Options) Offset() uint64 {
	return o.offset
}

func (o *Options) Fields() []Field {
	return o.fields
}

func (o *Options) AddField(name string, operator string, value string, fieldType string) error {
	err := validateOperator(operator)
	if err != nil {
		return err
	}

	if fieldType == DataTypeArray && operator != OperatorIn {
		return fmt.Errorf("with array type name you can use only `in` operator. wrong query param:`%s, %s, %s`",
			name, operator, value)
	}

	o.fields = append(o.fields, Field{
		Name:     name,
		Value:    value,
		Operator: operator,
		Type:     fieldType,
	})
	return nil
}

func validateOperator(operator string) error {
	switch operator {
	case
		OperatorEq,
		OperatorLike,
		OperatorNotEq,
		OperatorLowerThan,
		OperatorLowerThanEq,
		OperatorGreaterThan,
		OperatorGreaterThanEq,
		OperatorIn:
	default:
		return ErrBadOperator
	}
	return nil
}
