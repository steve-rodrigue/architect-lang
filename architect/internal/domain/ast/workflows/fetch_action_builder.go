package workflows

import "fmt"

type fetchActionBuilder struct {
	variable  TypedVariable
	source    string
	condition Condition
}

func (b *fetchActionBuilder) Variable(variable TypedVariable) FetchActionBuilder {
	b.variable = variable
	return b
}

func (b *fetchActionBuilder) Source(source string) FetchActionBuilder {
	b.source = source
	return b
}

func (b *fetchActionBuilder) Condition(condition Condition) FetchActionBuilder {
	b.condition = condition
	return b
}

func (b *fetchActionBuilder) Build() (FetchAction, error) {
	if b.variable == nil {
		return nil, fmt.Errorf("fetch action variable is required")
	}

	if b.source == "" {
		return nil, fmt.Errorf("fetch action source is required")
	}

	if b.condition == nil {
		return nil, fmt.Errorf("fetch action condition is required")
	}

	return &fetchAction{
		variable:  b.variable,
		source:    b.source,
		condition: b.condition,
	}, nil
}
