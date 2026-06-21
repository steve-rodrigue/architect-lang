package endpoints

import "fmt"

type inputSourceRuleBuilder struct {
	kind    InputSourceRuleKind
	sources []InputSource
}

func (b *inputSourceRuleBuilder) Kind(kind InputSourceRuleKind) InputSourceRuleBuilder {
	b.kind = kind

	return b
}

func (b *inputSourceRuleBuilder) AddSource(source InputSource) InputSourceRuleBuilder {
	if source != nil {
		b.sources = append(b.sources, source)
	}

	return b
}

func (b *inputSourceRuleBuilder) Build() (InputSourceRule, error) {
	switch b.kind {
	case InputSourceRuleSingle, InputSourceRuleExactlyOneOf:
		// valid
	default:
		return nil, fmt.Errorf("invalid input source rule kind %q", b.kind)
	}

	if len(b.sources) == 0 {
		return nil, fmt.Errorf("input source rule must contain at least one source")
	}

	switch b.kind {
	case InputSourceRuleSingle:
		if len(b.sources) != 1 {
			return nil, fmt.Errorf("single input source rule must contain exactly one source")
		}

	case InputSourceRuleExactlyOneOf:
		if len(b.sources) < 2 {
			return nil, fmt.Errorf("exactly_one_of input source rule must contain at least two sources")
		}
	}

	sources := make([]InputSource, len(b.sources))
	copy(sources, b.sources)

	return &inputSourceRule{
		kind:    b.kind,
		sources: sources,
	}, nil
}
