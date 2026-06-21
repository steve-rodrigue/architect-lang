package workflows

import "fmt"

type createActionBuilder struct {
	variable    TypedVariable
	assignments []Assignment
}

func (b *createActionBuilder) Variable(variable TypedVariable) CreateActionBuilder {
	b.variable = variable
	return b
}

func (b *createActionBuilder) AddAssignment(assignment Assignment) CreateActionBuilder {
	if assignment != nil {
		b.assignments = append(b.assignments, assignment)
	}

	return b
}

func (b *createActionBuilder) Build() (CreateAction, error) {
	if b.variable == nil {
		return nil, fmt.Errorf("create action variable is required")
	}

	assignments := make([]Assignment, len(b.assignments))
	copy(assignments, b.assignments)

	return &createAction{
		variable:    b.variable,
		assignments: assignments,
	}, nil
}
