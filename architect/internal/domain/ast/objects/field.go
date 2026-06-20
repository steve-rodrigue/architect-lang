package objects

type field struct {
	name      string
	typeRef   TypeReference
	modifiers []FieldModifier
}

func (f *field) Name() string {
	return f.name
}

func (f *field) Type() TypeReference {
	return f.typeRef
}

func (f *field) Modifiers() []FieldModifier {
	return append([]FieldModifier(nil), f.modifiers...)
}
