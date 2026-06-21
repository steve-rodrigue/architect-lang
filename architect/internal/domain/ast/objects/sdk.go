package objects

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"

type FieldModifierKind string

const (
	FieldModifierPrimary     FieldModifierKind = "primary"
	FieldModifierUnique      FieldModifierKind = "unique"
	FieldModifierRenamedFrom FieldModifierKind = "renamed_from"
	FieldModifierDeprecated  FieldModifierKind = "deprecated"
)

// NewObjectBuilder creates a new object builder
func NewObjectBuilder() ObjectBuilder {
	return &objectBuilder{}
}

// NewFieldBuilder creates a new field builder
func NewFieldBuilder() FieldBuilder {
	return &fieldBuilder{}
}

// NewFieldModifierBuilder creates a new field modifier builder
func NewFieldModifierBuilder() FieldModifierBuilder {
	return &fieldModifierBuilder{}
}

// ObjectBuilder represents an object builder
type ObjectBuilder interface {
	Name(name string) ObjectBuilder
	HistoryOf(typeRef common.TypeReference) ObjectBuilder
	AddField(field Field) ObjectBuilder
	Build() (Object, error)
}

// Object represents an object
type Object interface {
	Name() string
	HistoryOf() common.TypeReference
	Fields() []Field
}

// FieldBuilder represents a field builder
type FieldBuilder interface {
	Name(name string) FieldBuilder
	Type(typeRef common.TypeReference) FieldBuilder
	AddModifier(modifier FieldModifier) FieldBuilder
	Build() (Field, error)
}

// Field represents a field
type Field interface {
	Name() string
	Type() common.TypeReference
	Modifiers() []FieldModifier
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
