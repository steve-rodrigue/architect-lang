package objects

type typeReference struct {
	name             string
	isList           bool
	isOptional       bool
	numberConstraint NumberConstraint
	defaultValue     Value
}

func (t *typeReference) Name() string {
	return t.name
}

func (t *typeReference) IsList() bool {
	return t.isList
}

func (t *typeReference) IsOptional() bool {
	return t.isOptional
}

func (t *typeReference) NumberConstraint() NumberConstraint {
	return t.numberConstraint
}

func (t *typeReference) DefaultValue() Value {
	return t.defaultValue
}

func (t *typeReference) HasDefaultValue() bool {
	return t.defaultValue != nil
}
