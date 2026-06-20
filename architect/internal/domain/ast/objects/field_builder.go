package objects

import "fmt"

type fieldBuilder struct {
	name      string
	typeRef   TypeReference
	modifiers []FieldModifier
}

func (b *fieldBuilder) Name(name string) FieldBuilder {
	b.name = name

	return b
}

func (b *fieldBuilder) Type(typeRef TypeReference) FieldBuilder {
	b.typeRef = typeRef

	return b
}

func (b *fieldBuilder) AddModifier(modifier FieldModifier) FieldBuilder {
	if modifier != nil {
		b.modifiers = append(b.modifiers, modifier)
	}

	return b
}

func (b *fieldBuilder) Build() (Field, error) {
	if b.name == "" {
		return nil, fmt.Errorf("field name is required")
	}

	if b.typeRef == nil {
		return nil, fmt.Errorf("field %q type is required", b.name)
	}

	var hasPrimary bool
	var hasUnique bool
	var hasRenamedFrom bool
	var hasDeprecated bool

	for _, modifier := range b.modifiers {
		switch {
		case modifier.IsPrimary():
			if hasPrimary {
				return nil, fmt.Errorf("field %q has multiple primary modifiers", b.name)
			}

			hasPrimary = true

		case modifier.IsUnique():
			if hasUnique {
				return nil, fmt.Errorf("field %q has multiple unique modifiers", b.name)
			}

			hasUnique = true

		case modifier.IsRenameFrom():
			if hasRenamedFrom {
				return nil, fmt.Errorf("field %q has multiple renamed_from modifiers", b.name)
			}

			hasRenamedFrom = true

		case modifier.IsDeprecated():
			if hasDeprecated {
				return nil, fmt.Errorf("field %q has multiple deprecated modifiers", b.name)
			}

			hasDeprecated = true
		}
	}

	return &field{
		name:      b.name,
		typeRef:   b.typeRef,
		modifiers: append([]FieldModifier(nil), b.modifiers...),
	}, nil
}
