package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type assignment struct {
	name  string
	value Expression
	ast   workflows.Assignment
}

func (a *assignment) Name() string {
	return a.name
}

func (a *assignment) Value() Expression {
	return a.value
}

func (a *assignment) AST() workflows.Assignment {
	return a.ast
}
