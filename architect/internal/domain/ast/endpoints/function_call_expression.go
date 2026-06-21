package endpoints

type functionCallExpression struct {
	name      string
	arguments []Expression
}

func (e *functionCallExpression) Kind() ExpressionKind {
	return ExpressionKindFunctionCall
}

func (e *functionCallExpression) Name() string {
	return e.name
}

func (e *functionCallExpression) Arguments() []Expression {
	arguments := make([]Expression, len(e.arguments))
	copy(arguments, e.arguments)

	return arguments
}
