package endpoints

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
)

type literalExpressionBuilder struct {
	value objects.Value
}

func (b *literalExpressionBuilder) Value(value objects.Value) LiteralExpressionBuilder {
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
