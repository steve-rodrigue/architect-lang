package workflows

type condition struct {
	left     Expression
	operator Comparator
	right    Expression
}

func (c *condition) Left() Expression {
	return c.left
}

func (c *condition) Operator() Comparator {
	return c.operator
}

func (c *condition) Right() Expression {
	return c.right
}
