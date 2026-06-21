package endpoints

type inputSourceRule struct {
	kind    InputSourceRuleKind
	sources []InputSource
}

func (r *inputSourceRule) Kind() InputSourceRuleKind {
	return r.kind
}

func (r *inputSourceRule) Sources() []InputSource {
	sources := make([]InputSource, len(r.sources))
	copy(sources, r.sources)

	return sources
}
