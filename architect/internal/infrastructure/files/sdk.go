package files

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/applications/parsers"
	projects_application "github.com/steve-rodrigue/architect-lang/architect/internal/applications/projects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

// NewPlanService creates a new file-backend plan service
func NewPlanService(
	path string,
	fileName string,
) plan.Service {
	return &planService{
		path:     path,
		fileName: fileName,
	}
}

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
