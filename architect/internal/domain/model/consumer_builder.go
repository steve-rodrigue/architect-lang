package model

import "fmt"

type consumerBuilder struct {
	event        Event
	actions      []Action
	astEventName string
}

func (b *consumerBuilder) Event(event Event) ConsumerBuilder {
	b.event = event
	return b
}

func (b *consumerBuilder) AddAction(action Action) ConsumerBuilder {
	b.actions = append(b.actions, action)
	return b
}

func (b *consumerBuilder) ASTEventName(name string) ConsumerBuilder {
	b.astEventName = name
	return b
}

func (b *consumerBuilder) Build() (Consumer, error) {
	if b.event == nil {
		return nil, fmt.Errorf("consumer event is required")
	}

	for _, action := range b.actions {
		if action == nil {
			return nil, fmt.Errorf("consumer action is required")
		}
	}

	return &consumer{
		event:        b.event,
		actions:      append([]Action(nil), b.actions...),
		astEventName: b.astEventName,
	}, nil
}
