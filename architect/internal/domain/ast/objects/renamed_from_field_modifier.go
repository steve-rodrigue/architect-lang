package objects

type renamedFromFieldModifier struct {
	fieldModifier
	name string
}

func (m *renamedFromFieldModifier) Name() string {
	return m.name
}
