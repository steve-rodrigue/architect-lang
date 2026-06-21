package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type functionCallExpressionBuilder struct {
	name      string
	arguments []Expression
	ast       workflows.FunctionCallExpression
}

func (b *functionCallExpressionBuilder) Name(name string) FunctionCallExpressionBuilder {
	b.name = name
	return b
}

func (b *functionCallExpressionBuilder) AddArgument(argument Expression) FunctionCallExpressionBuilder {
	b.arguments = append(b.arguments, argument)
	return b
}

func (b *functionCallExpressionBuilder) AST(ast workflows.FunctionCallExpression) FunctionCallExpressionBuilder {
	b.ast = ast
	return b
}

func (b *functionCallExpressionBuilder) Build() (FunctionCallExpression, error) {
	if b.name == "" {
		return nil, fmt.Errorf("function call expression name is required")
	}

	for _, argument := range b.arguments {
		if argument == nil {
			return nil, fmt.Errorf("function call expression argument is required")
		}
	}

	return &functionCallExpression{
		name:      b.name,
		arguments: append([]Expression(nil), b.arguments...),
		ast:       b.ast,
	}, nil
}
