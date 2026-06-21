package endpoints

import "fmt"

type updateActionBuilder struct {
	variableName string
	assignments  []Assignment
}

func (b *updateActionBuilder) VariableName(name string) UpdateActionBuilder {
	b.variableName = name
	return b
}

func (b *updateActionBuilder) AddAssignment(assignment Assignment) UpdateActionBuilder {
	if assignment != nil {
		b.assignments = append(b.assignments, assignment)
	}

	return b
}

func (b *updateActionBuilder) Build() (UpdateAction, error) {
	if b.variableName == "" {
		return nil, fmt.Errorf("update action variable name is required")
	}

	assignments := make([]Assignment, len(b.assignments))
	copy(assignments, b.assignments)

	return &updateAction{
		variableName: b.variableName,
		assignments:  assignments,
	}, nil
}
