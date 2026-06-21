package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type fetchAction struct {
	variable  TypedVariable
	source    Service
	condition Condition
	ast       workflows.FetchAction
}

func (a *fetchAction) Kind() workflows.ActionKind {
	return workflows.ActionKindFetch
}

func (a *fetchAction) AST() workflows.Action {
	return a.ast
}

func (a *fetchAction) Variable() TypedVariable {
	return a.variable
}

func (a *fetchAction) Source() Service {
	return a.source
}

func (a *fetchAction) Condition() Condition {
	return a.condition
}
