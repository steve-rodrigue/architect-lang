package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
)

type inputBuilder struct {
	fields []InputField
	ast    endpoints.Input
}

func (b *inputBuilder) AddField(field InputField) InputBuilder {
	b.fields = append(b.fields, field)
	return b
}

func (b *inputBuilder) AST(ast endpoints.Input) InputBuilder {
	b.ast = ast
	return b
}

func (b *inputBuilder) Build() (Input, error) {
	for _, field := range b.fields {
		if field == nil {
			return nil, fmt.Errorf("input field is required")
		}
	}

	return &input{
		fields: append([]InputField(nil), b.fields...),
		ast:    b.ast,
	}, nil
}
