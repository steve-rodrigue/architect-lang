package workflows

type storeAction struct {
	variableName string
	destination  string
}

func (a *storeAction) Kind() ActionKind {
	return ActionKindStore
}

func (a *storeAction) VariableName() string {
	return a.variableName
}

func (a *storeAction) Destination() string {
	return a.destination
}
