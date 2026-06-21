package workflows

import "fmt"

type storeActionBuilder struct {
	variableName string
	destination  string
}

func (b *storeActionBuilder) VariableName(name string) StoreActionBuilder {
	b.variableName = name
	return b
}

func (b *storeActionBuilder) Destination(destination string) StoreActionBuilder {
	b.destination = destination
	return b
}

func (b *storeActionBuilder) Build() (StoreAction, error) {
	if b.variableName == "" {
		return nil, fmt.Errorf("store action variable name is required")
	}

	if b.destination == "" {
		return nil, fmt.Errorf("store action destination is required")
	}

	return &storeAction{
		variableName: b.variableName,
		destination:  b.destination,
	}, nil
}
