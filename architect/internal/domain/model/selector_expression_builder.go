package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

type selectorExpressionBuilder struct {
	parts []string
	ast   workflows.SelectorExpression
}

func (b *selectorExpressionBuilder) AddPart(part string) SelectorExpressionBuilder {
	b.parts = append(b.parts, part)
	return b
}

func (b *selectorExpressionBuilder) AST(ast workflows.SelectorExpression) SelectorExpressionBuilder {
	b.ast = ast
	return b
}

func (b *selectorExpressionBuilder) Build() (SelectorExpression, error) {
	if len(b.parts) == 0 {
		return nil, fmt.Errorf("selector expression parts are required")
	}

	for _, part := range b.parts {
		if part == "" {
			return nil, fmt.Errorf("selector expression part is required")
		}
	}

	return &selectorExpression{
		parts: append([]string(nil), b.parts...),
		ast:   b.ast,
	}, nil
}
