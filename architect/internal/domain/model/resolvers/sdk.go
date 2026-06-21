package resolvers

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model/sources"
)

// NewResolver creates a new model resolver.
func NewResolver() Resolver {
	return &resolver{}
}

// Resolver resolves a parsed project source into a linked project model.
type Resolver interface {
	Resolve(source sources.Project) (model.Project, error)
}
