package endpoints

import "fmt"

type selectorExpressionBuilder struct {
	parts []string
}

func (b *selectorExpressionBuilder) AddPart(part string) SelectorExpressionBuilder {
	b.parts = append(b.parts, part)

	return b
}

func (b *selectorExpressionBuilder) Build() (SelectorExpression, error) {
	if len(b.parts) == 0 {
		return nil, fmt.Errorf("selector expression must contain at least one part")
	}

	parts := make([]string, len(b.parts))
	copy(parts, b.parts)

	return &selectorExpression{
		parts: parts,
	}, nil
}
