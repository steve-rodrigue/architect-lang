package workflows

import "fmt"

type functionCallExpressionBuilder struct {
	name      string
	arguments []Expression
}

func (b *functionCallExpressionBuilder) Name(name string) FunctionCallExpressionBuilder {
	b.name = name

	return b
}

func (b *functionCallExpressionBuilder) AddArgument(argument Expression) FunctionCallExpressionBuilder {
	if argument != nil {
		b.arguments = append(b.arguments, argument)
	}

	return b
}

func (b *functionCallExpressionBuilder) Build() (FunctionCallExpression, error) {
	if b.name == "" {
		return nil, fmt.Errorf("function call expression name is required")
	}

	arguments := make([]Expression, len(b.arguments))
	copy(arguments, b.arguments)

	return &functionCallExpression{
		name:      b.name,
		arguments: arguments,
	}, nil
}
