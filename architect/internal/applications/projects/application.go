package projects

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model/sources"
)

// NewApplication creates a new project application.
func NewApplication(resolver model.Resolver) Application {
	return &application{
		resolver: resolver,
	}
}

type application struct {
	resolver model.Resolver
	project  model.Project
	loaded   bool
}

func (a *application) Load(source sources.Project) error {
	if a.resolver == nil {
		return fmt.Errorf("project resolver is required")
	}

	if source == nil {
		return fmt.Errorf("project source is required")
	}

	project, err := a.resolver.Resolve(source)
	if err != nil {
		return err
	}

	if project == nil {
		return fmt.Errorf("resolved project is nil")
	}

	a.project = project
	a.loaded = true

	return nil
}

func (a *application) Project() (model.Project, error) {
	if !a.loaded || a.project == nil {
		return nil, fmt.Errorf("project is not loaded")
	}

	return a.project, nil
}
