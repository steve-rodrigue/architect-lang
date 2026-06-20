package objects

type FieldModifierKind string

const (
	FieldModifierPrimary     FieldModifierKind = "primary"
	FieldModifierUnique      FieldModifierKind = "unique"
	FieldModifierRenamedFrom FieldModifierKind = "renamed_from"
	FieldModifierDeprecated  FieldModifierKind = "deprecated"
)

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

// NewObjectBuilder creates a new object builder
func NewObjectBuilder() ObjectBuilder {
	return &objectBuilder{}
}

// NewFieldBuilder creates a new field builder
func NewFieldBuilder() FieldBuilder {
	return &fieldBuilder{}
}

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

// NewFieldModifierBuilder creates a new field modifier builder
func NewFieldModifierBuilder() FieldModifierBuilder {
	return &fieldModifierBuilder{}
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return &valueBuilder{}
}

// ObjectBuilder represents an object builder
type ObjectBuilder interface {
	Name(name string) ObjectBuilder
	HistoryOf(typeRef TypeReference) ObjectBuilder
	AddField(field Field) ObjectBuilder
	Build() (Object, error)
}

// Object represents an object
type Object interface {
	Name() string
	HistoryOf() TypeReference
	Fields() []Field
}

// FieldBuilder represents a field builder
type FieldBuilder interface {
	Name(name string) FieldBuilder
	Type(typeRef TypeReference) FieldBuilder
	AddModifier(modifier FieldModifier) FieldBuilder
	Build() (Field, error)
}

// Field represents a field
type Field interface {
	Name() string
	Type() TypeReference
	Modifiers() []FieldModifier
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

// FieldModifierBuilder represents a field modifier builder
type FieldModifierBuilder interface {
	Primary() (FieldModifier, error)
	Unique() (FieldModifier, error)
	RenamedFrom(name string) (RenamedFromFieldModifier, error)
	Deprecated() (FieldModifier, error)
}

// FieldModifier represents a field modifier
type FieldModifier interface {
	Kind() FieldModifierKind
	IsPrimary() bool
	IsUnique() bool
	IsRenameFrom() bool
	IsDeprecated() bool
}

// RenamedFromFieldModifier represents a renamed from field modifier
type RenamedFromFieldModifier interface {
	FieldModifier
	Name() string
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
