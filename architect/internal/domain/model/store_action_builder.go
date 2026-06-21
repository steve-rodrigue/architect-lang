package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type storeActionBuilder struct {
	variableName string
	destination  Service
	ast          workflows.StoreAction
}

func (b *storeActionBuilder) VariableName(name string) StoreActionBuilder {
	b.variableName = name
	return b
}

func (b *storeActionBuilder) Destination(service Service) StoreActionBuilder {
	b.destination = service
	return b
}

func (b *storeActionBuilder) AST(ast workflows.StoreAction) StoreActionBuilder {
	b.ast = ast
	return b
}

func (b *storeActionBuilder) Build() (StoreAction, error) {
	if b.variableName == "" {
		return nil, fmt.Errorf("store action variable name is required")
	}

	if b.destination == nil {
		return nil, fmt.Errorf("store action destination service is required")
	}

	return &storeAction{
		variableName: b.variableName,
		destination:  b.destination,
		ast:          b.ast,
	}, nil
}
