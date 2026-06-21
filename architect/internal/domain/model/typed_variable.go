package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type typedVariable struct {
	name string
	typ  TypeReference
	ast  workflows.TypedVariable
}

func (v *typedVariable) Name() string {
	return v.name
}

func (v *typedVariable) Type() TypeReference {
	return v.typ
}

func (v *typedVariable) AST() workflows.TypedVariable {
	return v.ast
}
