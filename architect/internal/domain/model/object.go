package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"

type object struct {
	name      string
	historyOf Object
	fields    []Field
	ast       objects.Object
}

func (o *object) Name() string {
	return o.name
}

func (o *object) HistoryOf() Object {
	return o.historyOf
}

func (o *object) Fields() []Field {
	return append([]Field(nil), o.fields...)
}

func (o *object) AST() objects.Object {
	return o.ast
}
