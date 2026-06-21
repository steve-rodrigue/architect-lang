package common

type ValueKind string

const (
	ValueString ValueKind = "string"
	ValueInt    ValueKind = "int"
	ValueFloat  ValueKind = "float"
	ValueBool   ValueKind = "bool"
)

type NumberConstraintKind string

const (
	NumberConstraintNone       NumberConstraintKind = "none"
	NumberConstraintOneOrMore  NumberConstraintKind = "one_or_more"
	NumberConstraintZeroOrMore NumberConstraintKind = "zero_or_more"
	NumberConstraintRange      NumberConstraintKind = "range"
)

// NewTypeReferenceBuilder creates a new type reference builder
func NewTypeReferenceBuilder() TypeReferenceBuilder {
	return &typeReferenceBuilder{}
}

// NewNumberConstraintBuilder creates a new number contraint builder
func NewNumberConstraintBuilder() NumberConstraintBuilder {
	return &numberConstraintBuilder{}
}

// NewNumberValueBuilder creates a new number value builder
func NewNumberValueBuilder() NumberValueBuilder {
	return &numberValueBuilder{}
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return &valueBuilder{}
}

// TypeReferenceBuilder represents a type reference builder
type TypeReferenceBuilder interface {
	Name(name string) TypeReferenceBuilder
	List(isList bool) TypeReferenceBuilder
	Optional(isOptional bool) TypeReferenceBuilder
	NumberConstraint(constraint NumberConstraint) TypeReferenceBuilder
	DefaultValue(value Value) TypeReferenceBuilder
	Build() (TypeReference, error)
}

// TypeReference represents a type reference
type TypeReference interface {
	Name() string
	IsList() bool
	IsOptional() bool

	NumberConstraint() NumberConstraint

	DefaultValue() Value
	HasDefaultValue() bool
}

// NumberConstraintBuilder represents a number constraint builder
type NumberConstraintBuilder interface {
	Kind(kind NumberConstraintKind) NumberConstraintBuilder
	Min(value NumberValue) NumberConstraintBuilder
	Max(value NumberValue) NumberConstraintBuilder
	Build() (NumberConstraint, error)
}

// NumberConstraint represents a number constraint
type NumberConstraint interface {
	Kind() NumberConstraintKind

	Min() NumberValue
	Max() NumberValue

	HasMin() bool
	HasMax() bool
}

// NumberValueBuilder represents a number value builder
type NumberValueBuilder interface {
	Int(raw string) (NumberValue, error)
	Float(raw string) (NumberValue, error)
}

// NumberValue represents a number value
type NumberValue interface {
	Raw() string
	IsFloat() bool
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	String(raw string) (Value, error)
	Int(raw string) (Value, error)
	Float(raw string) (Value, error)
	Bool(raw string) (Value, error)
}

// Value represents a value
type Value interface {
	Kind() ValueKind
	Raw() string
}
