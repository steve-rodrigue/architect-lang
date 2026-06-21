package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type conditionBuilder struct {
	left     Expression
	operator workflows.Comparator
	right    Expression
	ast      workflows.Condition
}

func (b *conditionBuilder) Left(expression Expression) ConditionBuilder {
	b.left = expression
	return b
}

func (b *conditionBuilder) Operator(operator workflows.Comparator) ConditionBuilder {
	b.operator = operator
	return b
}

func (b *conditionBuilder) Right(expression Expression) ConditionBuilder {
	b.right = expression
	return b
}

func (b *conditionBuilder) AST(ast workflows.Condition) ConditionBuilder {
	b.ast = ast
	return b
}

func (b *conditionBuilder) Build() (Condition, error) {
	if b.left == nil {
		return nil, fmt.Errorf("condition left expression is required")
	}

	if b.right == nil {
		return nil, fmt.Errorf("condition right expression is required")
	}

	switch b.operator {
	case workflows.ComparatorEqual,
		workflows.ComparatorNotEqual,
		workflows.ComparatorLessThan,
		workflows.ComparatorLessThanOrEqual,
		workflows.ComparatorGreaterThan,
		workflows.ComparatorGreaterThanOrEqual:
	default:
		return nil, fmt.Errorf("condition operator is required")
	}

	return &condition{
		left:     b.left,
		operator: b.operator,
		right:    b.right,
		ast:      b.ast,
	}, nil
}
