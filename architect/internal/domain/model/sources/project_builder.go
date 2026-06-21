package sources

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/projects"
)

type projectBuilder struct {
	project  projects.Project
	versions []Version
}

func (b *projectBuilder) Project(project projects.Project) ProjectBuilder {
	b.project = project
	return b
}

func (b *projectBuilder) AddVersion(version Version) ProjectBuilder {
	b.versions = append(b.versions, version)
	return b
}

func (b *projectBuilder) Build() (Project, error) {
	if b.project == nil {
		return nil, fmt.Errorf("project source project is required")
	}

	for _, version := range b.versions {
		if version == nil {
			return nil, fmt.Errorf("project source version is required")
		}
	}

	return &project{
		project:  b.project,
		versions: append([]Version(nil), b.versions...),
	}, nil
}
