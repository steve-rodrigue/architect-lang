package workflows

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
)

type literalExpressionBuilder struct {
	value common.Value
}

func (b *literalExpressionBuilder) Value(value common.Value) LiteralExpressionBuilder {
	b.value = value

	return b
}

func (b *literalExpressionBuilder) Build() (LiteralExpression, error) {
	if b.value == nil {
		return nil, fmt.Errorf("literal expression value is required")
	}

	return &literalExpression{
		value: b.value,
	}, nil
}
