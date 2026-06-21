package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
)

type inputFieldBuilder struct {
	name    string
	typ     TypeReference
	sources endpoints.InputSourceRule
	ast     endpoints.InputField
}

func (b *inputFieldBuilder) Name(name string) InputFieldBuilder {
	b.name = name
	return b
}

func (b *inputFieldBuilder) Type(typeRef TypeReference) InputFieldBuilder {
	b.typ = typeRef
	return b
}

func (b *inputFieldBuilder) Sources(rule endpoints.InputSourceRule) InputFieldBuilder {
	b.sources = rule
	return b
}

func (b *inputFieldBuilder) AST(ast endpoints.InputField) InputFieldBuilder {
	b.ast = ast
	return b
}

func (b *inputFieldBuilder) Build() (InputField, error) {
	if b.name == "" {
		return nil, fmt.Errorf("input field name is required")
	}

	if b.typ == nil {
		return nil, fmt.Errorf("input field type is required")
	}

	if b.sources == nil {
		return nil, fmt.Errorf("input field sources are required")
	}

	return &inputField{
		name:    b.name,
		typ:     b.typ,
		sources: b.sources,
		ast:     b.ast,
	}, nil
}
