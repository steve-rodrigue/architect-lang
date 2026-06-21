package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"

type inputField struct {
	name    string
	typ     TypeReference
	sources endpoints.InputSourceRule
	ast     endpoints.InputField
}

func (f *inputField) Name() string {
	return f.name
}

func (f *inputField) Type() TypeReference {
	return f.typ
}

func (f *inputField) Sources() endpoints.InputSourceRule {
	return f.sources
}

func (f *inputField) AST() endpoints.InputField {
	return f.ast
}
