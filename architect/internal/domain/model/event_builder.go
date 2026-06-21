package model

import "fmt"

type eventBuilder struct {
	name       string
	declaredBy Service
}

func (b *eventBuilder) Name(name string) EventBuilder {
	b.name = name
	return b
}

func (b *eventBuilder) DeclaredBy(service Service) EventBuilder {
	b.declaredBy = service
	return b
}

func (b *eventBuilder) Build() (Event, error) {
	if b.name == "" {
		return nil, fmt.Errorf("event name is required")
	}

	return &event{
		name:       b.name,
		declaredBy: b.declaredBy,
	}, nil
}
