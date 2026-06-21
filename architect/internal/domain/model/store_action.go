package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type storeAction struct {
	variableName string
	destination  Service
	ast          workflows.StoreAction
}

func (a *storeAction) Kind() workflows.ActionKind {
	return workflows.ActionKindStore
}

func (a *storeAction) AST() workflows.Action {
	return a.ast
}

func (a *storeAction) VariableName() string {
	return a.variableName
}

func (a *storeAction) Destination() Service {
	return a.destination
}
