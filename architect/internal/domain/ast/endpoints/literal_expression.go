package endpoints

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"

type literalExpression struct {
	value objects.Value
}

func (e *literalExpression) Kind() ExpressionKind {
	return ExpressionKindLiteral
}

func (e *literalExpression) Value() objects.Value {
	return e.value
}
