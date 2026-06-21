package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type callActionBuilder struct {
	targetService   Service
	targetOperation string
	assignments     []Assignment
	ast             workflows.CallAction
}

func (b *callActionBuilder) TargetService(service Service) CallActionBuilder {
	b.targetService = service
	return b
}

func (b *callActionBuilder) TargetOperation(operation string) CallActionBuilder {
	b.targetOperation = operation
	return b
}

func (b *callActionBuilder) AddAssignment(assignment Assignment) CallActionBuilder {
	b.assignments = append(b.assignments, assignment)
	return b
}

func (b *callActionBuilder) AST(ast workflows.CallAction) CallActionBuilder {
	b.ast = ast
	return b
}

func (b *callActionBuilder) Build() (CallAction, error) {
	if b.targetService == nil {
		return nil, fmt.Errorf("call action target service is required")
	}

	if b.targetOperation == "" {
		return nil, fmt.Errorf("call action target operation is required")
	}

	for _, assignment := range b.assignments {
		if assignment == nil {
			return nil, fmt.Errorf("call action assignment is required")
		}
	}

	return &callAction{
		targetService:   b.targetService,
		targetOperation: b.targetOperation,
		assignments:     append([]Assignment(nil), b.assignments...),
		ast:             b.ast,
	}, nil
}
