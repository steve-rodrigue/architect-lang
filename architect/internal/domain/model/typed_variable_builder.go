package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type typedVariableBuilder struct {
	name string
	typ  TypeReference
	ast  workflows.TypedVariable
}

func (b *typedVariableBuilder) Name(name string) TypedVariableBuilder {
	b.name = name
	return b
}

func (b *typedVariableBuilder) Type(typeRef TypeReference) TypedVariableBuilder {
	b.typ = typeRef
	return b
}

func (b *typedVariableBuilder) AST(ast workflows.TypedVariable) TypedVariableBuilder {
	b.ast = ast
	return b
}

func (b *typedVariableBuilder) Build() (TypedVariable, error) {
	if b.name == "" {
		return nil, fmt.Errorf("typed variable name is required")
	}

	if b.typ == nil {
		return nil, fmt.Errorf("typed variable type is required")
	}

	return &typedVariable{
		name: b.name,
		typ:  b.typ,
		ast:  b.ast,
	}, nil
}
