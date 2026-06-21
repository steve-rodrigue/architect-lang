package endpoints

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"

type inputField struct {
	name    string
	typeRef objects.TypeReference
	sources InputSourceRule
}

func (f *inputField) Name() string {
	return f.name
}

func (f *inputField) Type() objects.TypeReference {
	return f.typeRef
}

func (f *inputField) Sources() InputSourceRule {
	return f.sources
}
