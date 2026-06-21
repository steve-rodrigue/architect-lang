package endpoints

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
)

type inputField struct {
	name    string
	typeRef common.TypeReference
	sources InputSourceRule
}

func (f *inputField) Name() string {
	return f.name
}

func (f *inputField) Type() common.TypeReference {
	return f.typeRef
}

func (f *inputField) Sources() InputSourceRule {
	return f.sources
}
