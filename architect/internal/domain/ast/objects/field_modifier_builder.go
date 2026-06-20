package objects

import "fmt"

type fieldModifierBuilder struct{}

func (b *fieldModifierBuilder) Primary() (FieldModifier, error) {
	return &fieldModifier{kind: FieldModifierPrimary}, nil
}

func (b *fieldModifierBuilder) Unique() (FieldModifier, error) {
	return &fieldModifier{kind: FieldModifierUnique}, nil
}

func (b *fieldModifierBuilder) Deprecated() (FieldModifier, error) {
	return &fieldModifier{kind: FieldModifierDeprecated}, nil
}

func (b *fieldModifierBuilder) RenamedFrom(name string) (RenamedFromFieldModifier, error) {
	if name == "" {
		return nil, fmt.Errorf("renamed_from name cannot be empty")
	}

	return &renamedFromFieldModifier{
		fieldModifier: fieldModifier{kind: FieldModifierRenamedFrom},
		name:          name,
	}, nil
}
