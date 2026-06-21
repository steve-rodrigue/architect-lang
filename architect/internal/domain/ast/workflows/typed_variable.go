package workflows

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
)

type typedVariable struct {
	name    string
	typeRef common.TypeReference
}

func (v *typedVariable) Name() string {
	return v.name
}

func (v *typedVariable) Type() common.TypeReference {
	return v.typeRef
}
