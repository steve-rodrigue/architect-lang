package endpoints

import "fmt"

type returnActionBuilder struct {
	expression Expression
}

func (b *returnActionBuilder) Expression(expression Expression) ReturnActionBuilder {
	b.expression = expression
	return b
}

func (b *returnActionBuilder) Build() (ReturnAction, error) {
	if b.expression == nil {
		return nil, fmt.Errorf("return action expression is required")
	}

	return &returnAction{expression: b.expression}, nil
}
