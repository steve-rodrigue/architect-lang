package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type returnAction struct {
	expression Expression
	ast        workflows.ReturnAction
}

func (a *returnAction) Kind() workflows.ActionKind {
	return workflows.ActionKindReturn
}

func (a *returnAction) AST() workflows.Action {
	return a.ast
}

func (a *returnAction) Expression() Expression {
	return a.expression
}
