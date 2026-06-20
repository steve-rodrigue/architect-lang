package objects

type numberConstraint struct {
	kind NumberConstraintKind
	min  NumberValue
	max  NumberValue
}

func (c *numberConstraint) Kind() NumberConstraintKind {
	return c.kind
}

func (c *numberConstraint) Min() NumberValue {
	return c.min
}

func (c *numberConstraint) Max() NumberValue {
	return c.max
}

func (c *numberConstraint) HasMin() bool {
	return c.min != nil
}

func (c *numberConstraint) HasMax() bool {
	return c.max != nil
}
