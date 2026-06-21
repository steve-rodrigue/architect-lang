package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type selectorExpression struct {
	parts []string
	ast   workflows.SelectorExpression
}

func (e *selectorExpression) Kind() workflows.ExpressionKind {
	return workflows.ExpressionKindSelector
}

func (e *selectorExpression) AST() workflows.Expression {
	return e.ast
}

func (e *selectorExpression) Parts() []string {
	return append([]string(nil), e.parts...)
}
