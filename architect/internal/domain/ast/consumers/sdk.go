package consumers

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

// NewConsumerBuilder creates a new consumer builder
func NewConsumerBuilder() ConsumerBuilder {
	return &consumerBuilder{
		actions: make([]workflows.Action, 0),
	}
}

// ConsumerBuilder represents a consumer builder
type ConsumerBuilder interface {
	EventName(name string) ConsumerBuilder
	AddAction(action workflows.Action) ConsumerBuilder
	Build() (Consumer, error)
}

// Consumer represents an event consumer
type Consumer interface {
	EventName() string
	Actions() []workflows.Action
}
