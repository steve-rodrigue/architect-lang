package endpoints

type fetchAction struct {
	variable  TypedVariable
	source    string
	condition Condition
}

func (a *fetchAction) Kind() ActionKind {
	return ActionKindFetch
}

func (a *fetchAction) Variable() TypedVariable {
	return a.variable
}

func (a *fetchAction) Source() string {
	return a.source
}

func (a *fetchAction) Condition() Condition {
	return a.condition
}
