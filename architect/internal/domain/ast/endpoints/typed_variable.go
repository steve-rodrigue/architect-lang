package endpoints

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"

type typedVariable struct {
	name    string
	typeRef objects.TypeReference
}

func (v *typedVariable) Name() string {
	return v.name
}

func (v *typedVariable) Type() objects.TypeReference {
	return v.typeRef
}
