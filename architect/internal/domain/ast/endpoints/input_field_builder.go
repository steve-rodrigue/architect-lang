package endpoints

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
)

type inputFieldBuilder struct {
	name    string
	typeRef objects.TypeReference
	sources InputSourceRule
}

func (b *inputFieldBuilder) Name(name string) InputFieldBuilder {
	b.name = name
	return b
}

func (b *inputFieldBuilder) Type(typeRef objects.TypeReference) InputFieldBuilder {
	b.typeRef = typeRef
	return b
}

func (b *inputFieldBuilder) Sources(rule InputSourceRule) InputFieldBuilder {
	b.sources = rule
	return b
}

func (b *inputFieldBuilder) Build() (InputField, error) {
	if b.name == "" {
		return nil, fmt.Errorf("input field name is required")
	}

	if b.typeRef == nil {
		return nil, fmt.Errorf("input field %q type is required", b.name)
	}

	if b.sources == nil {
		return nil, fmt.Errorf("input field %q sources are required", b.name)
	}

	return &inputField{
		name:    b.name,
		typeRef: b.typeRef,
		sources: b.sources,
	}, nil
}
