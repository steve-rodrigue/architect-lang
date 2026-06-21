package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"

type typeReference struct {
	name             string
	object           Object
	isList           bool
	isOptional       bool
	numberConstraint common.NumberConstraint
	defaultValue     common.Value
	ast              common.TypeReference
}

func (t *typeReference) Name() string {
	return t.name
}

func (t *typeReference) Object() Object {
	return t.object
}

func (t *typeReference) IsObject() bool {
	return t.object != nil
}

func (t *typeReference) IsList() bool {
	return t.isList
}

func (t *typeReference) IsOptional() bool {
	return t.isOptional
}

func (t *typeReference) NumberConstraint() common.NumberConstraint {
	return t.numberConstraint
}

func (t *typeReference) DefaultValue() common.Value {
	return t.defaultValue
}

func (t *typeReference) HasDefaultValue() bool {
	return t.defaultValue != nil
}

func (t *typeReference) AST() common.TypeReference {
	return t.ast
}
