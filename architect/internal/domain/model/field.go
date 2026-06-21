package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"

type field struct {
	name      string
	typ       TypeReference
	modifiers []objects.FieldModifier
	ast       objects.Field
}

func (f *field) Name() string {
	return f.name
}

func (f *field) Type() TypeReference {
	return f.typ
}

func (f *field) Modifiers() []objects.FieldModifier {
	return append([]objects.FieldModifier(nil), f.modifiers...)
}

func (f *field) AST() objects.Field {
	return f.ast
}
