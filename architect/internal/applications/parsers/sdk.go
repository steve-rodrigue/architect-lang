package parsers

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/consumers"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
)

// Application represents the parser application
type Application interface {
	Object(script string) (objects.Object, error)
	Endpoint(script string) (endpoints.Endpoint, error)
	Consumer(script string) (consumers.Consumer, error)
	Application(script string) (applications.Application, error)
}
