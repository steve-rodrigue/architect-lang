package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"

type endpoint struct {
	name    string
	method  endpoints.HTTPMethod
	path    string
	input   Input
	actions []Action
	ast     endpoints.Endpoint
}

func (e *endpoint) Name() string {
	return e.name
}

func (e *endpoint) Method() endpoints.HTTPMethod {
	return e.method
}

func (e *endpoint) Path() string {
	return e.path
}

func (e *endpoint) Input() Input {
	return e.input
}

func (e *endpoint) Actions() []Action {
	return append([]Action(nil), e.actions...)
}

func (e *endpoint) AST() endpoints.Endpoint {
	return e.ast
}
