package model

type consumer struct {
	event        Event
	actions      []Action
	astEventName string
}

func (c *consumer) Event() Event {
	return c.event
}

func (c *consumer) Actions() []Action {
	return append([]Action(nil), c.actions...)
}

func (c *consumer) ASTEventName() string {
	return c.astEventName
}
