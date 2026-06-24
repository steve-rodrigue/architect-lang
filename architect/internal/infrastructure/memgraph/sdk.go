package memgraph

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

// NewPlanService creates a new plan service
func NewPlanService(
	driver neo4j.DriverWithContext,
	database string,
) plan.Service {
	return &planService{
		driver:   driver,
		database: database,
	}
}

// NewPlanProjectRepository creates a new plan project repository
func NewPlanProjectRepository(
	driver neo4j.DriverWithContext,
	database string,
) plan.ProjectRepository {
	return &planProjectRepository{
		driver:   driver,
		database: database,
	}
}

// NewPlanVersionRepository creates a new plan version repository
func NewPlanVersionRepository(
	driver neo4j.DriverWithContext,
	database string,
	projectRepository plan.ProjectRepository,
) plan.VersionRepository {
	return &planVersionRepository{
		driver:            driver,
		database:          database,
		projectRepository: projectRepository,
	}
}

// NewPlanVersionRepository creates a new plan section repository
func NewPlanSectionRepository(
	driver neo4j.DriverWithContext,
	database string,
	versionRepository plan.VersionRepository,
) plan.SectionRepository {
	return &planSectionRepository{
		driver:            driver,
		database:          database,
		versionRepository: versionRepository,
	}
}

// NewPlanDependencyRepository creates a new plan dependency repository
func NewPlanDependencyRepository(
	driver neo4j.DriverWithContext,
	database string,
	versionRepository plan.VersionRepository,
) plan.DependencyRepository {
	return &planDependencyRepository{
		driver:            driver,
		database:          database,
		versionRepository: versionRepository,
	}
}

// NewPlanArtifactRepository creates a new plan artifact repository
func NewPlanArtifactRepository(
	driver neo4j.DriverWithContext,
	database string,
	sectionRepository plan.SectionRepository,
) plan.ArtifactRepository {
	return &planArtifactRepository{
		driver:            driver,
		database:          database,
		sectionRepository: sectionRepository,
	}
}

// NewPlanTaskRepository creates a new plan task repository
func NewPlanTaskRepository(
	driver neo4j.DriverWithContext,
	database string,
	sectionRepository plan.SectionRepository,
) plan.TaskRepository {
	return &planTaskRepository{
		driver:            driver,
		database:          database,
		sectionRepository: sectionRepository,
	}
}
