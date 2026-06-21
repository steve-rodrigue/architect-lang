package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"

type condition struct {
	left     Expression
	operator workflows.Comparator
	right    Expression
	ast      workflows.Condition
}

func (c *condition) Left() Expression {
	return c.left
}

func (c *condition) Operator() workflows.Comparator {
	return c.operator
}

func (c *condition) Right() Expression {
	return c.right
}

func (c *condition) AST() workflows.Condition {
	return c.ast
}
