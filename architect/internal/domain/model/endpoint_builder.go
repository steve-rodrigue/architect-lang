package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
)

type endpointBuilder struct {
	name    string
	method  endpoints.HTTPMethod
	path    string
	input   Input
	actions []Action
	ast     endpoints.Endpoint
}

func (b *endpointBuilder) Name(name string) EndpointBuilder {
	b.name = name
	return b
}

func (b *endpointBuilder) Method(method endpoints.HTTPMethod) EndpointBuilder {
	b.method = method
	return b
}

func (b *endpointBuilder) Path(path string) EndpointBuilder {
	b.path = path
	return b
}

func (b *endpointBuilder) Input(input Input) EndpointBuilder {
	b.input = input
	return b
}

func (b *endpointBuilder) AddAction(action Action) EndpointBuilder {
	b.actions = append(b.actions, action)
	return b
}

func (b *endpointBuilder) AST(ast endpoints.Endpoint) EndpointBuilder {
	b.ast = ast
	return b
}

func (b *endpointBuilder) Build() (Endpoint, error) {
	if b.name == "" {
		return nil, fmt.Errorf("endpoint name is required")
	}

	if b.method == "" {
		return nil, fmt.Errorf("endpoint method is required")
	}

	if b.path == "" {
		return nil, fmt.Errorf("endpoint path is required")
	}

	if b.input == nil {
		return nil, fmt.Errorf("endpoint input is required")
	}

	for _, action := range b.actions {
		if action == nil {
			return nil, fmt.Errorf("endpoint action is required")
		}
	}

	return &endpoint{
		name:    b.name,
		method:  b.method,
		path:    b.path,
		input:   b.input,
		actions: append([]Action(nil), b.actions...),
		ast:     b.ast,
	}, nil
}
