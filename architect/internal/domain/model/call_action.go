package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type callAction struct {
	targetService   Service
	targetOperation string
	assignments     []Assignment
	ast             workflows.CallAction
}

func (a *callAction) Kind() workflows.ActionKind {
	return workflows.ActionKindCall
}

func (a *callAction) AST() workflows.Action {
	return a.ast
}

func (a *callAction) TargetService() Service {
	return a.targetService
}

func (a *callAction) TargetOperation() string {
	return a.targetOperation
}

func (a *callAction) Assignments() []Assignment {
	return append([]Assignment(nil), a.assignments...)
}
