package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type updateActionBuilder struct {
	variableName string
	assignments  []Assignment
	ast          workflows.UpdateAction
}

func (b *updateActionBuilder) VariableName(name string) UpdateActionBuilder {
	b.variableName = name
	return b
}

func (b *updateActionBuilder) AddAssignment(assignment Assignment) UpdateActionBuilder {
	b.assignments = append(b.assignments, assignment)
	return b
}

func (b *updateActionBuilder) AST(ast workflows.UpdateAction) UpdateActionBuilder {
	b.ast = ast
	return b
}

func (b *updateActionBuilder) Build() (UpdateAction, error) {
	if b.variableName == "" {
		return nil, fmt.Errorf("update action variable name is required")
	}

	for _, assignment := range b.assignments {
		if assignment == nil {
			return nil, fmt.Errorf("update action assignment is required")
		}
	}

	return &updateAction{
		variableName: b.variableName,
		assignments:  append([]Assignment(nil), b.assignments...),
		ast:          b.ast,
	}, nil
}
