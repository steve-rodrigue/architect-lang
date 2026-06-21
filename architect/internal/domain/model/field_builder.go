package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
)

type fieldBuilder struct {
	name      string
	typ       TypeReference
	modifiers []objects.FieldModifier
	ast       objects.Field
}

func (b *fieldBuilder) Name(name string) FieldBuilder {
	b.name = name
	return b
}

func (b *fieldBuilder) Type(typeRef TypeReference) FieldBuilder {
	b.typ = typeRef
	return b
}

func (b *fieldBuilder) AddModifier(modifier objects.FieldModifier) FieldBuilder {
	b.modifiers = append(b.modifiers, modifier)
	return b
}

func (b *fieldBuilder) AST(ast objects.Field) FieldBuilder {
	b.ast = ast
	return b
}

func (b *fieldBuilder) Build() (Field, error) {
	if b.name == "" {
		return nil, fmt.Errorf("field name is required")
	}

	if b.typ == nil {
		return nil, fmt.Errorf("field type is required")
	}

	for _, modifier := range b.modifiers {
		if modifier == nil {
			return nil, fmt.Errorf("field modifier is required")
		}
	}

	return &field{
		name:      b.name,
		typ:       b.typ,
		modifiers: append([]objects.FieldModifier(nil), b.modifiers...),
		ast:       b.ast,
	}, nil
}
