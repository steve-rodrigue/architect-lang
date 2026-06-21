package endpoints

import "fmt"

type endpointBuilder struct {
	name    string
	method  HTTPMethod
	path    string
	input   Input
	actions []Action
}

func (b *endpointBuilder) Name(name string) EndpointBuilder {
	b.name = name

	return b
}

func (b *endpointBuilder) Method(method HTTPMethod) EndpointBuilder {
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
	if action != nil {
		b.actions = append(b.actions, action)
	}

	return b
}

func (b *endpointBuilder) Build() (Endpoint, error) {
	if b.name == "" {
		return nil, fmt.Errorf("endpoint name is required")
	}

	switch b.method {
	case HTTPMethodGet,
		HTTPMethodPost,
		HTTPMethodPut,
		HTTPMethodPatch,
		HTTPMethodDelete:
		// valid

	default:
		return nil, fmt.Errorf("invalid endpoint method %q", b.method)
	}

	if b.path == "" {
		return nil, fmt.Errorf("endpoint %q path is required", b.name)
	}

	if b.input == nil {
		return nil, fmt.Errorf("endpoint %q input is required", b.name)
	}

	actions := make([]Action, len(b.actions))
	copy(actions, b.actions)

	return &endpoint{
		name:    b.name,
		method:  b.method,
		path:    b.path,
		input:   b.input,
		actions: actions,
	}, nil
}
