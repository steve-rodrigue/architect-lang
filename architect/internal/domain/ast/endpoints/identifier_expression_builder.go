package endpoints

import "fmt"

type identifierExpressionBuilder struct {
	name string
}

func (b *identifierExpressionBuilder) Name(name string) IdentifierExpressionBuilder {
	b.name = name

	return b
}

func (b *identifierExpressionBuilder) Build() (IdentifierExpression, error) {
	if b.name == "" {
		return nil, fmt.Errorf("identifier expression name is required")
	}

	return &identifierExpression{
		name: b.name,
	}, nil
}
