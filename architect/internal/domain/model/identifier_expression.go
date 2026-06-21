package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type identifierExpression struct {
	name string
	ast  workflows.IdentifierExpression
}

func (e *identifierExpression) Kind() workflows.ExpressionKind {
	return workflows.ExpressionKindIdentifier
}

func (e *identifierExpression) AST() workflows.Expression {
	return e.ast
}

func (e *identifierExpression) Name() string {
	return e.name
}
