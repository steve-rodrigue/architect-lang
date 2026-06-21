package consumers

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type consumerBuilder struct {
	eventName string
	actions   []workflows.Action
}

func (b *consumerBuilder) EventName(name string) ConsumerBuilder {
	b.eventName = name

	return b
}

func (b *consumerBuilder) AddAction(action workflows.Action) ConsumerBuilder {
	b.actions = append(b.actions, action)

	return b
}

func (b *consumerBuilder) Build() (Consumer, error) {
	if b.eventName == "" {
		return nil, fmt.Errorf("event name is required")
	}

	if len(b.actions) == 0 {
		return nil, fmt.Errorf("at least one action is required")
	}

	actions := make([]workflows.Action, len(b.actions))
	copy(actions, b.actions)

	return &consumer{
		eventName: b.eventName,
		actions:   actions,
	}, nil
}
