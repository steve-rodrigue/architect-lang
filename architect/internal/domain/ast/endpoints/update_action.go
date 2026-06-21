package endpoints

type updateAction struct {
	variableName string
	assignments  []Assignment
}

func (a *updateAction) Kind() ActionKind {
	return ActionKindUpdate
}

func (a *updateAction) VariableName() string {
	return a.variableName
}

func (a *updateAction) Assignments() []Assignment {
	assignments := make([]Assignment, len(a.assignments))
	copy(assignments, a.assignments)

	return assignments
}
