package endpoints

type returnAction struct {
	expression Expression
}

func (a *returnAction) Kind() ActionKind {
	return ActionKindReturn
}

func (a *returnAction) Expression() Expression {
	return a.expression
}
