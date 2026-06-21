package endpoints

type input struct {
	fields []InputField
}

func (i *input) Fields() []InputField {
	fields := make([]InputField, len(i.fields))
	copy(fields, i.fields)

	return fields
}
