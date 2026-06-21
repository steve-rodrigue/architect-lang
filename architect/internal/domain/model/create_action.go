package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type createAction struct {
	variable    TypedVariable
	assignments []Assignment
	ast         workflows.CreateAction
}

func (a *createAction) Kind() workflows.ActionKind {
	return workflows.ActionKindCreate
}

func (a *createAction) AST() workflows.Action {
	return a.ast
}

func (a *createAction) Variable() TypedVariable {
	return a.variable
}

func (a *createAction) Assignments() []Assignment {
	return append([]Assignment(nil), a.assignments...)
}
