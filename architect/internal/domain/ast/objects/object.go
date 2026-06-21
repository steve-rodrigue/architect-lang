package objects

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"

type object struct {
	name      string
	historyOf common.TypeReference
	fields    []Field
}

func (o *object) Name() string {
	return o.name
}

func (o *object) HistoryOf() common.TypeReference {
	return o.historyOf
}

func (o *object) Fields() []Field {
	return append([]Field(nil), o.fields...)
}
