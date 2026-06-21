package endpoints

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
)

type typedVariableBuilder struct {
	name    string
	typeRef objects.TypeReference
}

func (b *typedVariableBuilder) Name(name string) TypedVariableBuilder {
	b.name = name
	return b
}

func (b *typedVariableBuilder) Type(typeRef objects.TypeReference) TypedVariableBuilder {
	b.typeRef = typeRef
	return b
}

func (b *typedVariableBuilder) Build() (TypedVariable, error) {
	if b.name == "" {
		return nil, fmt.Errorf("typed variable name is required")
	}

	if b.typeRef == nil {
		return nil, fmt.Errorf("typed variable %q type is required", b.name)
	}

	return &typedVariable{name: b.name, typeRef: b.typeRef}, nil
}
