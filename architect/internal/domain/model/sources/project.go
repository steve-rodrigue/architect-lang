package sources

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/projects"

type project struct {
	project  projects.Project
	versions []Version
}

func (s *project) Project() projects.Project {
	return s.project
}

func (s *project) Versions() []Version {
	return append([]Version(nil), s.versions...)
}
