package workflows

import "fmt"

type callActionBuilder struct {
	target      string
	assignments []Assignment
}

func (b *callActionBuilder) Target(target string) CallActionBuilder {
	b.target = target
	return b
}

func (b *callActionBuilder) AddAssignment(assignment Assignment) CallActionBuilder {
	if assignment != nil {
		b.assignments = append(b.assignments, assignment)
	}

	return b
}

func (b *callActionBuilder) Build() (CallAction, error) {
	if b.target == "" {
		return nil, fmt.Errorf("call action target is required")
	}

	assignments := make([]Assignment, len(b.assignments))
	copy(assignments, b.assignments)

	return &callAction{
		target:      b.target,
		assignments: assignments,
	}, nil
}
