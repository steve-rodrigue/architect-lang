package endpoints

import "fmt"

type assignmentBuilder struct {
	name  string
	value Expression
}

func (b *assignmentBuilder) Name(name string) AssignmentBuilder {
	b.name = name

	return b
}

func (b *assignmentBuilder) Value(value Expression) AssignmentBuilder {
	b.value = value

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
	}, nil
}
