package files

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/applications/parsers"
	projects_application "github.com/steve-rodrigue/architect-lang/architect/internal/applications/projects"
)

// NewProjectApplication creates a new file-backed project application.
func NewProjectApplication(
	parserApplication parsers.Application,
	projectApp projects_application.Application,
) projects_application.FileApplication {
	return &projectApplication{
		parserApplication:  parserApplication,
		projectApplication: projectApp,
	}
}
