package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type emitActionBuilder struct {
	event       Event
	variable    TypedVariable
	assignments []Assignment
	ast         workflows.EmitAction
}

func (b *emitActionBuilder) Event(event Event) EmitActionBuilder {
	b.event = event
	return b
}

func (b *emitActionBuilder) Variable(variable TypedVariable) EmitActionBuilder {
	b.variable = variable
	return b
}

func (b *emitActionBuilder) AddAssignment(assignment Assignment) EmitActionBuilder {
	b.assignments = append(b.assignments, assignment)
	return b
}

func (b *emitActionBuilder) AST(ast workflows.EmitAction) EmitActionBuilder {
	b.ast = ast
	return b
}

func (b *emitActionBuilder) Build() (EmitAction, error) {
	if b.event == nil {
		return nil, fmt.Errorf("emit action event is required")
	}

	if b.variable == nil {
		return nil, fmt.Errorf("emit action variable is required")
	}

	for _, assignment := range b.assignments {
		if assignment == nil {
			return nil, fmt.Errorf("emit action assignment is required")
		}
	}

	return &emitAction{
		event:       b.event,
		variable:    b.variable,
		assignments: append([]Assignment(nil), b.assignments...),
		ast:         b.ast,
	}, nil
}
