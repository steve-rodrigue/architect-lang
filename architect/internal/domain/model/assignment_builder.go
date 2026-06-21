package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type assignmentBuilder struct {
	name  string
	value Expression
	ast   workflows.Assignment
}

func (b *assignmentBuilder) Name(name string) AssignmentBuilder {
	b.name = name
	return b
}

func (b *assignmentBuilder) Value(value Expression) AssignmentBuilder {
	b.value = value
	return b
}

func (b *assignmentBuilder) AST(ast workflows.Assignment) AssignmentBuilder {
	b.ast = ast
	return b
}

func (b *assignmentBuilder) Build() (Assignment, error) {
	if b.name == "" {
		return nil, fmt.Errorf("assignment name is required")
	}

	if b.value == nil {
		return nil, fmt.Errorf("assignment value is required")
	}

	return &assignment{
		name:  b.name,
		value: b.value,
		ast:   b.ast,
	}, nil
}
