package projects

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model/sources"
)

// FileApplication represents a file application
type FileApplication interface {
	Parse(relativePath string) error
	Load() error
	Project() (model.Project, error)
}

// Application loads and resolves a project model.
type Application interface {
	Load(source sources.Project) error
	Project() (model.Project, error)
}
