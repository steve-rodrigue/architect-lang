package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type createActionBuilder struct {
	variable    TypedVariable
	assignments []Assignment
	ast         workflows.CreateAction
}

func (b *createActionBuilder) Variable(variable TypedVariable) CreateActionBuilder {
	b.variable = variable
	return b
}

func (b *createActionBuilder) AddAssignment(assignment Assignment) CreateActionBuilder {
	b.assignments = append(b.assignments, assignment)
	return b
}

func (b *createActionBuilder) AST(ast workflows.CreateAction) CreateActionBuilder {
	b.ast = ast
	return b
}

func (b *createActionBuilder) Build() (CreateAction, error) {
	if b.variable == nil {
		return nil, fmt.Errorf("create action variable is required")
	}

	for _, assignment := range b.assignments {
		if assignment == nil {
			return nil, fmt.Errorf("create action assignment is required")
		}
	}

	return &createAction{
		variable:    b.variable,
		assignments: append([]Assignment(nil), b.assignments...),
		ast:         b.ast,
	}, nil
}
