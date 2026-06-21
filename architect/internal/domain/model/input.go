package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"

type input struct {
	fields []InputField
	ast    endpoints.Input
}

func (i *input) Fields() []InputField {
	return append([]InputField(nil), i.fields...)
}

func (i *input) AST() endpoints.Input {
	return i.ast
}
