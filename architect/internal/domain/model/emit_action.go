package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type emitAction struct {
	event       Event
	variable    TypedVariable
	assignments []Assignment
	ast         workflows.EmitAction
}

func (a *emitAction) Kind() workflows.ActionKind {
	return workflows.ActionKindEmit
}

func (a *emitAction) AST() workflows.Action {
	return a.ast
}

func (a *emitAction) Event() Event {
	return a.event
}

func (a *emitAction) Variable() TypedVariable {
	return a.variable
}

func (a *emitAction) Assignments() []Assignment {
	return append([]Assignment(nil), a.assignments...)
}
