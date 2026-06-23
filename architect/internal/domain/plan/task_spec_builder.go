package plan

import "fmt"

type taskSpecBuilder struct {
	goal               string
	instructions       string
	constraints        []string
	acceptanceCriteria []string
}

func (b *taskSpecBuilder) Goal(goal string) TaskSpecBuilder {
	b.goal = goal
	return b
}

func (b *taskSpecBuilder) Instructions(instructions string) TaskSpecBuilder {
	b.instructions = instructions
	return b
}

func (b *taskSpecBuilder) AddConstraint(constraint string) TaskSpecBuilder {
	b.constraints = append(b.constraints, constraint)
	return b
}

func (b *taskSpecBuilder) AddAcceptanceCriteria(criteria string) TaskSpecBuilder {
	b.acceptanceCriteria = append(b.acceptanceCriteria, criteria)
	return b
}

func (b *taskSpecBuilder) Build() (TaskSpec, error) {
	if b.goal == "" {
		return nil, fmt.Errorf("task spec goal is required")
	}

	if b.instructions == "" {
		return nil, fmt.Errorf("task spec instructions are required")
	}

	return &taskSpec{
		goal:               b.goal,
		instructions:       b.instructions,
		constraints:        append([]string(nil), b.constraints...),
		acceptanceCriteria: append([]string(nil), b.acceptanceCriteria...),
	}, nil
}
