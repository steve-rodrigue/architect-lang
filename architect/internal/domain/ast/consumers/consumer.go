package consumers

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type consumer struct {
	eventName string
	actions   []workflows.Action
}

func (c *consumer) EventName() string {
	return c.eventName
}

func (c *consumer) Actions() []workflows.Action {
	actions := make([]workflows.Action, len(c.actions))
	copy(actions, c.actions)

	return actions
}
