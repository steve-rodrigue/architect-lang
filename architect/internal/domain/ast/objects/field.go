package objects

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"

type field struct {
	name      string
	typeRef   common.TypeReference
	modifiers []FieldModifier
}

func (f *field) Name() string {
	return f.name
}

func (f *field) Type() common.TypeReference {
	return f.typeRef
}

func (f *field) Modifiers() []FieldModifier {
	return append([]FieldModifier(nil), f.modifiers...)
}
