package plan

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"

type project struct {
	id            ProjectID
	model         model.Project
	schemaVersion string
}

func (p *project) ID() ProjectID {
	return p.id
}

func (p *project) Model() model.Project {
	return p.model
}

func (p *project) SchemaVersion() string {
	return p.schemaVersion
}
