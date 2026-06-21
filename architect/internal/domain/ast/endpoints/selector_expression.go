package endpoints

type selectorExpression struct {
	parts []string
}

func (e *selectorExpression) Kind() ExpressionKind {
	return ExpressionKindSelector
}

func (e *selectorExpression) Parts() []string {
	parts := make([]string, len(e.parts))
	copy(parts, e.parts)

	return parts
}
