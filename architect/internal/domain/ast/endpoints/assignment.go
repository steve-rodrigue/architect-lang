package endpoints

type assignment struct {
	name  string
	value Expression
}

func (a *assignment) Name() string {
	return a.name
}

func (a *assignment) Value() Expression {
	return a.value
}
