package workflows

type callAction struct {
	target      string
	assignments []Assignment
}

func (a *callAction) Kind() ActionKind {
	return ActionKindCall
}

func (a *callAction) Target() string {
	return a.target
}

func (a *callAction) Assignments() []Assignment {
	assignments := make([]Assignment, len(a.assignments))
	copy(assignments, a.assignments)

	return assignments
}
