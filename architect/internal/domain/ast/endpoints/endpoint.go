package endpoints

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type endpoint struct {
	name    string
	method  HTTPMethod
	path    string
	input   Input
	actions []workflows.Action
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

func (e *endpoint) Actions() []workflows.Action {
	actions := make([]workflows.Action, len(e.actions))
	copy(actions, e.actions)

	return actions
}
