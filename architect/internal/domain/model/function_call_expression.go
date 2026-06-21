package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type functionCallExpression struct {
	name      string
	arguments []Expression
	ast       workflows.FunctionCallExpression
}

func (e *functionCallExpression) Kind() workflows.ExpressionKind {
	return workflows.ExpressionKindFunctionCall
}

func (e *functionCallExpression) AST() workflows.Expression {
	return e.ast
}

func (e *functionCallExpression) Name() string {
	return e.name
}

func (e *functionCallExpression) Arguments() []Expression {
	return append([]Expression(nil), e.arguments...)
}
