package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type literalExpressionBuilder struct {
	value common.Value
	ast   workflows.LiteralExpression
}

func (b *literalExpressionBuilder) Value(value common.Value) LiteralExpressionBuilder {
	b.value = value
	return b
}

func (b *literalExpressionBuilder) AST(ast workflows.LiteralExpression) LiteralExpressionBuilder {
	b.ast = ast
	return b
}

func (b *literalExpressionBuilder) Build() (LiteralExpression, error) {
	if b.value == nil {
		return nil, fmt.Errorf("literal expression value is required")
	}

	return &literalExpression{
		value: b.value,
		ast:   b.ast,
	}, nil
}
