package objects

import "fmt"

type numberConstraintBuilder struct {
	kind NumberConstraintKind
	min  NumberValue
	max  NumberValue
}

func (b *numberConstraintBuilder) Kind(kind NumberConstraintKind) NumberConstraintBuilder {
	b.kind = kind

	return b
}

func (b *numberConstraintBuilder) Min(value NumberValue) NumberConstraintBuilder {
	b.min = value

	return b
}

func (b *numberConstraintBuilder) Max(value NumberValue) NumberConstraintBuilder {
	b.max = value

	return b
}

func (b *numberConstraintBuilder) Build() (NumberConstraint, error) {
	if b.kind == "" {
		return nil, fmt.Errorf("number constraint kind is required")
	}

	switch b.kind {
	case NumberConstraintNone:
		if b.min != nil || b.max != nil {
			return nil, fmt.Errorf("none constraint cannot define min or max")
		}

	case NumberConstraintOneOrMore, NumberConstraintZeroOrMore:
		if b.min != nil || b.max != nil {
			return nil, fmt.Errorf("%s constraint cannot define min or max", b.kind)
		}

	case NumberConstraintRange:
		if b.min == nil && b.max == nil {
			return nil, fmt.Errorf("range constraint must define at least a min or a max")
		}

	default:
		return nil, fmt.Errorf("unknown number constraint kind %q", b.kind)
	}

	return &numberConstraint{
		kind: b.kind,
		min:  b.min,
		max:  b.max,
	}, nil
}
