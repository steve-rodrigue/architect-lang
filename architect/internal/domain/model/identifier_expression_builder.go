package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type identifierExpressionBuilder struct {
	name string
	ast  workflows.IdentifierExpression
}

func (b *identifierExpressionBuilder) Name(name string) IdentifierExpressionBuilder {
	b.name = name
	return b
}

func (b *identifierExpressionBuilder) AST(ast workflows.IdentifierExpression) IdentifierExpressionBuilder {
	b.ast = ast
	return b
}

func (b *identifierExpressionBuilder) Build() (IdentifierExpression, error) {
	if b.name == "" {
		return nil, fmt.Errorf("identifier expression name is required")
	}

	return &identifierExpression{
		name: b.name,
		ast:  b.ast,
	}, nil
}
