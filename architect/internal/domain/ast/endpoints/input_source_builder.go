package endpoints

import "fmt"

type inputSourceBuilder struct {
	kind     InputSourceKind
	optional bool
}

func (b *inputSourceBuilder) Kind(kind InputSourceKind) InputSourceBuilder {
	b.kind = kind

	return b
}

func (b *inputSourceBuilder) Optional(optional bool) InputSourceBuilder {
	b.optional = optional

	return b
}

func (b *inputSourceBuilder) Build() (InputSource, error) {
	switch b.kind {
	case InputSourcePath,
		InputSourceQuery,
		InputSourceBody,
		InputSourceHeader,
		InputSourceCookie:
		// valid

	default:
		return nil, fmt.Errorf("invalid input source kind %q", b.kind)
	}

	return &inputSource{
		kind:     b.kind,
		optional: b.optional,
	}, nil
}
