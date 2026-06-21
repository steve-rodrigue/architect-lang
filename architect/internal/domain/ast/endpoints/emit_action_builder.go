package endpoints

import "fmt"

type emitActionBuilder struct {
	variable    TypedVariable
	assignments []Assignment
}

func (b *emitActionBuilder) Variable(variable TypedVariable) EmitActionBuilder {
	b.variable = variable
	return b
}

func (b *emitActionBuilder) AddAssignment(assignment Assignment) EmitActionBuilder {
	if assignment != nil {
		b.assignments = append(b.assignments, assignment)
	}

	return b
}

func (b *emitActionBuilder) Build() (EmitAction, error) {
	if b.variable == nil {
		return nil, fmt.Errorf("emit action variable is required")
	}

	assignments := make([]Assignment, len(b.assignments))
	copy(assignments, b.assignments)

	return &emitAction{
		variable:    b.variable,
		assignments: assignments,
	}, nil
}
