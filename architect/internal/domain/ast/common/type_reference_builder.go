package common

import "fmt"

type typeReferenceBuilder struct {
	name             string
	isList           bool
	isOptional       bool
	numberConstraint NumberConstraint
	defaultValue     Value
}

func (b *typeReferenceBuilder) Name(name string) TypeReferenceBuilder {
	b.name = name
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

func (b *typeReferenceBuilder) NumberConstraint(constraint NumberConstraint) TypeReferenceBuilder {
	b.numberConstraint = constraint
	return b
}

func (b *typeReferenceBuilder) DefaultValue(value Value) TypeReferenceBuilder {
	b.defaultValue = value
	return b
}

func (b *typeReferenceBuilder) Build() (TypeReference, error) {
	if b.name == "" {
		return nil, fmt.Errorf("type reference name is required")
	}

	if b.defaultValue != nil && !b.isOptional {
		return nil, fmt.Errorf("default value can only be set on optional type %q", b.name)
	}

	return &typeReference{
		name:             b.name,
		isList:           b.isList,
		isOptional:       b.isOptional,
		numberConstraint: b.numberConstraint,
		defaultValue:     b.defaultValue,
	}, nil
}
