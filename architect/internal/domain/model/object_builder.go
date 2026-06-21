package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
)

type objectBuilder struct {
	name      string
	historyOf Object
	fields    []Field
	ast       objects.Object
}

func (b *objectBuilder) Name(name string) ObjectBuilder {
	b.name = name
	return b
}

func (b *objectBuilder) HistoryOf(object Object) ObjectBuilder {
	b.historyOf = object
	return b
}

func (b *objectBuilder) AddField(field Field) ObjectBuilder {
	b.fields = append(b.fields, field)
	return b
}

func (b *objectBuilder) AST(ast objects.Object) ObjectBuilder {
	b.ast = ast
	return b
}

func (b *objectBuilder) Build() (Object, error) {
	if b.name == "" {
		return nil, fmt.Errorf("object name is required")
	}

	for _, field := range b.fields {
		if field == nil {
			return nil, fmt.Errorf("object field is required")
		}
	}

	return &object{
		name:      b.name,
		historyOf: b.historyOf,
		fields:    append([]Field(nil), b.fields...),
		ast:       b.ast,
	}, nil
}
