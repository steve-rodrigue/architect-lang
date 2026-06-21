package endpoints

type emitAction struct {
	variable    TypedVariable
	assignments []Assignment
}

func (a *emitAction) Kind() ActionKind {
	return ActionKindEmit
}

func (a *emitAction) Variable() TypedVariable {
	return a.variable
}

func (a *emitAction) Assignments() []Assignment {
	assignments := make([]Assignment, len(a.assignments))
	copy(assignments, a.assignments)

	return assignments
}
