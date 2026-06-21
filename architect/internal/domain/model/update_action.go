package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type updateAction struct {
	variableName string
	assignments  []Assignment
	ast          workflows.UpdateAction
}

func (a *updateAction) Kind() workflows.ActionKind {
	return workflows.ActionKindUpdate
}

func (a *updateAction) AST() workflows.Action {
	return a.ast
}

func (a *updateAction) VariableName() string {
	return a.variableName
}

func (a *updateAction) Assignments() []Assignment {
	return append([]Assignment(nil), a.assignments...)
}
