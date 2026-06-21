package workflows

type identifierExpression struct {
	name string
}

func (e *identifierExpression) Kind() ExpressionKind {
	return ExpressionKindIdentifier
}

func (e *identifierExpression) Name() string {
	return e.name
}
