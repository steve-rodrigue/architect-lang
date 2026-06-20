package parsers

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"

// Application represents the parser application
type Application interface {
	Object(script string) (objects.Object, error)
}
