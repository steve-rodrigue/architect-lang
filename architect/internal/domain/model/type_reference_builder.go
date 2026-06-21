package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
)

type typeReferenceBuilder struct {
	name             string
	object           Object
	isList           bool
	isOptional       bool
	numberConstraint common.NumberConstraint
	defaultValue     common.Value
	ast              common.TypeReference
}

func (b *typeReferenceBuilder) Name(name string) TypeReferenceBuilder {
	b.name = name
	return b
}

func (b *typeReferenceBuilder) Object(object Object) TypeReferenceBuilder {
	b.object = object
	return b
}

func (b *typeReferenceBuilder) List(isList bool) TypeReferenceBuilder {
	b.isList = isList
	return b
}

func (b *typeReferenceBuilder) Optional(isOptional bool) TypeReferenceBuilder {
	b.isOptional = isOptional
	return b
}

func (b *typeReferenceBuilder) NumberConstraint(constraint common.NumberConstraint) TypeReferenceBuilder {
	b.numberConstraint = constraint
	return b
}

func (b *typeReferenceBuilder) DefaultValue(value common.Value) TypeReferenceBuilder {
	b.defaultValue = value
	return b
}

func (b *typeReferenceBuilder) AST(ast common.TypeReference) TypeReferenceBuilder {
	b.ast = ast
	return b
}

func (b *typeReferenceBuilder) Build() (TypeReference, error) {
	if b.name == "" {
		return nil, fmt.Errorf("type reference name is required")
	}

	return &typeReference{
		name:             b.name,
		object:           b.object,
		isList:           b.isList,
		isOptional:       b.isOptional,
		numberConstraint: b.numberConstraint,
		defaultValue:     b.defaultValue,
		ast:              b.ast,
	}, nil
}
