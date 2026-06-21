package endpoints

import "fmt"

type inputBuilder struct {
	fields []InputField
}

func (b *inputBuilder) AddField(field InputField) InputBuilder {
	if field != nil {
		b.fields = append(b.fields, field)
	}

	return b
}

func (b *inputBuilder) Build() (Input, error) {
	if len(b.fields) == 0 {
		return nil, fmt.Errorf("input must contain at least one field")
	}

	fields := make([]InputField, len(b.fields))
	copy(fields, b.fields)

	return &input{
		fields: fields,
	}, nil
}
