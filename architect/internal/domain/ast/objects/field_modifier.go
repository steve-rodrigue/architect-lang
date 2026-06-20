package objects

type fieldModifier struct {
	kind FieldModifierKind
}

func (m *fieldModifier) Kind() FieldModifierKind {
	return m.kind
}

func (m *fieldModifier) IsPrimary() bool {
	return m.kind == FieldModifierPrimary
}

func (m *fieldModifier) IsUnique() bool {
	return m.kind == FieldModifierUnique
}

func (m *fieldModifier) IsRenameFrom() bool {
	return m.kind == FieldModifierRenamedFrom
}

func (m *fieldModifier) IsDeprecated() bool {
	return m.kind == FieldModifierDeprecated
}
