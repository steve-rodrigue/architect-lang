package endpoints

type endpoint struct {
	name    string
	method  HTTPMethod
	path    string
	input   Input
	actions []Action
}

func (e *endpoint) Name() string {
	return e.name
}

func (e *endpoint) Method() HTTPMethod {
	return e.method
}

func (e *endpoint) Path() string {
	return e.path
}

func (e *endpoint) Input() Input {
	return e.input
}

func (e *endpoint) Actions() []Action {
	actions := make([]Action, len(e.actions))
	copy(actions, e.actions)

	return actions
}
