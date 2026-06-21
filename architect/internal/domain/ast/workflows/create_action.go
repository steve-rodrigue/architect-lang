package workflows

type createAction struct {
	variable    TypedVariable
	assignments []Assignment
}

func (a *createAction) Kind() ActionKind {
	return ActionKindCreate
}

func (a *createAction) Variable() TypedVariable {
	return a.variable
}

func (a *createAction) Assignments() []Assignment {
	assignments := make([]Assignment, len(a.assignments))
	copy(assignments, a.assignments)

	return assignments
}
