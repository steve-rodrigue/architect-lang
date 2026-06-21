package workflows

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
)

type literalExpression struct {
	value common.Value
}

func (e *literalExpression) Kind() ExpressionKind {
	return ExpressionKindLiteral
}

func (e *literalExpression) Value() common.Value {
	return e.value
}
