package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type returnActionBuilder struct {
	expression Expression
	ast        workflows.ReturnAction
}

func (b *returnActionBuilder) Expression(expression Expression) ReturnActionBuilder {
	b.expression = expression
	return b
}

func (b *returnActionBuilder) AST(ast workflows.ReturnAction) ReturnActionBuilder {
	b.ast = ast
	return b
}

func (b *returnActionBuilder) Build() (ReturnAction, error) {
	if b.expression == nil {
		return nil, fmt.Errorf("return action expression is required")
	}

	return &returnAction{
		expression: b.expression,
		ast:        b.ast,
	}, nil
}
