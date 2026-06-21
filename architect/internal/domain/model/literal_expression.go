package model

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type literalExpression struct {
	value common.Value
	ast   workflows.LiteralExpression
}

func (e *literalExpression) Kind() workflows.ExpressionKind {
	return workflows.ExpressionKindLiteral
}

func (e *literalExpression) AST() workflows.Expression {
	return e.ast
}

func (e *literalExpression) Value() common.Value {
	return e.value
}
