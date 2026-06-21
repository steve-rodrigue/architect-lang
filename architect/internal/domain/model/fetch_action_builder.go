package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type fetchActionBuilder struct {
	variable  TypedVariable
	source    Service
	condition Condition
	ast       workflows.FetchAction
}

func (b *fetchActionBuilder) Variable(variable TypedVariable) FetchActionBuilder {
	b.variable = variable
	return b
}

func (b *fetchActionBuilder) Source(service Service) FetchActionBuilder {
	b.source = service
	return b
}

func (b *fetchActionBuilder) Condition(condition Condition) FetchActionBuilder {
	b.condition = condition
	return b
}

func (b *fetchActionBuilder) AST(ast workflows.FetchAction) FetchActionBuilder {
	b.ast = ast
	return b
}

func (b *fetchActionBuilder) Build() (FetchAction, error) {
	if b.variable == nil {
		return nil, fmt.Errorf("fetch action variable is required")
	}

	if b.source == nil {
		return nil, fmt.Errorf("fetch action source service is required")
	}

	if b.condition == nil {
		return nil, fmt.Errorf("fetch action condition is required")
	}

	return &fetchAction{
		variable:  b.variable,
		source:    b.source,
		condition: b.condition,
		ast:       b.ast,
	}, nil
}
