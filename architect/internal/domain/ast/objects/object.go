package objects

type object struct {
	name      string
	historyOf TypeReference
	fields    []Field
}

func (o *object) Name() string {
	return o.name
}

func (o *object) HistoryOf() TypeReference {
	return o.historyOf
}

func (o *object) Fields() []Field {
	return append([]Field(nil), o.fields...)
}
