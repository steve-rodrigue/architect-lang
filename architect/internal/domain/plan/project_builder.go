package plan

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
)

type projectBuilder struct {
	model         model.Project
	schemaVersion string
}

func (b *projectBuilder) Model(projectModel model.Project) ProjectBuilder {
	b.model = projectModel
	return b
}

func (b *projectBuilder) SchemaVersion(schemaVersion string) ProjectBuilder {
	b.schemaVersion = schemaVersion
	return b
}

func (b *projectBuilder) Build() (Project, error) {
	if b.model == nil {
		return nil, fmt.Errorf("project model is required")
	}

	if b.model.Name() == "" {
		return nil, fmt.Errorf("project model name is required")
	}

	if b.schemaVersion == "" {
		return nil, fmt.Errorf("project schema version is required")
	}

	return &project{
		id:            GenProjectID(b.model.Name()),
		model:         b.model,
		schemaVersion: b.schemaVersion,
	}, nil
}
