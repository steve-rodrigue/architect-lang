package objects

import "fmt"

type objectBuilder struct {
	name      string
	historyOf TypeReference
	fields    []Field
}

func (b *objectBuilder) Name(name string) ObjectBuilder {
	b.name = name
	return b
}

func (b *objectBuilder) HistoryOf(typeRef TypeReference) ObjectBuilder {
	b.historyOf = typeRef
	return b
}

func (b *objectBuilder) AddField(field Field) ObjectBuilder {
	if field != nil {
		b.fields = append(b.fields, field)
	}
	return b
}

func (b *objectBuilder) Build() (Object, error) {
	if b.name == "" {
		return nil, fmt.Errorf("object name is required")
	}

	if len(b.fields) == 0 {
		return nil, fmt.Errorf("object %q must have at least one field", b.name)
	}

	seenFields := map[string]struct{}{}

	for _, field := range b.fields {
		if field == nil {
			return nil, fmt.Errorf("object %q contains nil field", b.name)
		}

		if _, exists := seenFields[field.Name()]; exists {
			return nil, fmt.Errorf("object %q has duplicate field %q", b.name, field.Name())
		}

		seenFields[field.Name()] = struct{}{}
	}

	return &object{
		name:      b.name,
		historyOf: b.historyOf,
		fields:    append([]Field(nil), b.fields...),
	}, nil
}
