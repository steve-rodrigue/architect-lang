package workflows

import "fmt"

type conditionBuilder struct {
	left     Expression
	operator Comparator
	right    Expression
}

func (b *conditionBuilder) Left(expression Expression) ConditionBuilder {
	b.left = expression
	return b
}

func (b *conditionBuilder) Operator(operator Comparator) ConditionBuilder {
	b.operator = operator
	return b
}

func (b *conditionBuilder) Right(expression Expression) ConditionBuilder {
	b.right = expression
	return b
}

func (b *conditionBuilder) Build() (Condition, error) {
	if b.left == nil {
		return nil, fmt.Errorf("condition left expression is required")
	}

	if b.operator == "" {
		return nil, fmt.Errorf("condition operator is required")
	}

	if b.right == nil {
		return nil, fmt.Errorf("condition right expression is required")
	}

	return &condition{
		left:     b.left,
		operator: b.operator,
		right:    b.right,
	}, nil
}
