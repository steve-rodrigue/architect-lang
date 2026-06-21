package workflows

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
)

type typedVariableBuilder struct {
	name    string
	typeRef common.TypeReference
}

func (b *typedVariableBuilder) Name(name string) TypedVariableBuilder {
	b.name = name
	return b
}

func (b *typedVariableBuilder) Type(typeRef common.TypeReference) TypedVariableBuilder {
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
