package memgraph

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

type planService struct {
	driver   neo4j.DriverWithContext
	database string
}

func (s *planService) SaveProject(ctx context.Context, project plan.Project) error {
	if project == nil {
		return fmt.Errorf("project is required")
	}

	return s.write(ctx, func(ctx context.Context, tx neo4j.ManagedTransaction) error {
		return run(ctx, tx, `
			MERGE (p:Project {id: $id})
			SET
				p.name = $name,
				p.schemaVersion = $schemaVersion
		`, map[string]any{
			"id":            string(project.ID()),
			"name":          project.Model().Name(),
			"schemaVersion": project.SchemaVersion(),
		})
	})
}

func (s *planService) SaveVersion(ctx context.Context, version plan.Version) error {
	if version == nil {
		return fmt.Errorf("version is required")
	}

	if version.Project() == nil {
		return fmt.Errorf("version project is required")
	}

	exists, err := s.hasProject(ctx, version.Project().ID())
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("version references unknown project %q", version.Project().ID())
	}

	return s.write(ctx, func(ctx context.Context, tx neo4j.ManagedTransaction) error {
		return run(ctx, tx, `
			MATCH (p:Project {id: $projectID})
			MERGE (v:Version {id: $id})
			SET
				v.number = $number
			MERGE (p)-[:HAS_VERSION]->(v)
		`, map[string]any{
			"id":        string(version.ID()),
			"projectID": string(version.Project().ID()),
			"number":    version.Number(),
		})
	})
}

func (s *planService) SaveSection(ctx context.Context, section plan.Section) error {
	if section == nil {
		return fmt.Errorf("section is required")
	}

	if section.Version() == nil {
		return fmt.Errorf("section version is required")
	}

	exists, err := s.hasVersion(ctx, section.Version().ID())
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("section references unknown version %q", section.Version().ID())
	}

	if section.HasParentID() {
		exists, err := s.hasSection(ctx, section.ParentID())
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("section references unknown parent section %q", section.ParentID())
		}
	}

	return s.write(ctx, func(ctx context.Context, tx neo4j.ManagedTransaction) error {
		if err := run(ctx, tx, `
			MATCH (v:Version {id: $versionID})
			MERGE (s:Section {id: $id})
			SET
				s.kind = $kind,
				s.name = $name,
				s.path = $path,
				s.parentID = $parentID
			MERGE (v)-[:HAS_SECTION]->(s)
		`, map[string]any{
			"id":        string(section.ID()),
			"versionID": string(section.Version().ID()),
			"kind":      string(section.Kind()),
			"name":      section.Name(),
			"path":      section.Path(),
			"parentID":  string(section.ParentID()),
		}); err != nil {
			return err
		}

		if section.HasParentID() {
			return run(ctx, tx, `
				MATCH
					(parent:Section {id: $parentID}),
					(child:Section {id: $childID})
				MERGE (parent)-[:PARENT_OF]->(child)
			`, map[string]any{
				"parentID": string(section.ParentID()),
				"childID":  string(section.ID()),
			})
		}

		return nil
	})
}

func (s *planService) SaveDependency(ctx context.Context, dependency plan.Dependency) error {
	if dependency == nil {
		return fmt.Errorf("dependency is required")
	}

	if dependency.Version() == nil {
		return fmt.Errorf("dependency version is required")
	}

	exists, err := s.hasVersion(ctx, dependency.Version().ID())
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("dependency references unknown version %q", dependency.Version().ID())
	}

	return s.write(ctx, func(ctx context.Context, tx neo4j.ManagedTransaction) error {
		return run(ctx, tx, `
			MATCH (v:Version {id: $versionID})
			MERGE (d:Dependency {id: $id})
			SET
				d.kind = $kind,
				d.name = $name,
				d.depVersion = $depVersion,
				d.source = $source
			MERGE (v)-[:HAS_DEPENDENCY]->(d)
		`, map[string]any{
			"id":         string(dependency.ID()),
			"versionID":  string(dependency.Version().ID()),
			"kind":       string(dependency.Kind()),
			"name":       dependency.Name(),
			"depVersion": dependency.DepVersion(),
			"source":     dependency.Source(),
		})
	})
}

func (s *planService) SaveArtifact(ctx context.Context, artifact plan.Artifact) error {
	if artifact == nil {
		return fmt.Errorf("artifact is required")
	}

	if artifact.Section() == nil {
		return fmt.Errorf("artifact section is required")
	}

	exists, err := s.hasSection(ctx, artifact.Section().ID())
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("artifact references unknown section %q", artifact.Section().ID())
	}

	return s.write(ctx, func(ctx context.Context, tx neo4j.ManagedTransaction) error {
		if err := run(ctx, tx, `
			MATCH (s:Section {id: $sectionID})
			MERGE (a:Artifact {id: $id})
			SET
				a.kind = $kind,
				a.path = $path,
				a.isFile = $isFile
			MERGE (s)-[:HAS_ARTIFACT]->(a)
		`, map[string]any{
			"id":        string(artifact.ID()),
			"sectionID": string(artifact.Section().ID()),
			"kind":      string(artifact.Kind()),
			"path":      artifact.Path(),
			"isFile":    artifact.IsFile(),
		}); err != nil {
			return err
		}

		if artifact.IsFile() && artifact.File() != nil {
			file := artifact.File()

			return run(ctx, tx, `
				MATCH (a:Artifact {id: $artifactID})
				MERGE (f:ArtifactFile {artifactID: $artifactID})
				SET
					f.kind = $kind,
					f.purpose = $purpose,
					f.specification = $specification,
					f.constraints = $constraints,
					f.acceptanceCriteria = $acceptanceCriteria
				MERGE (a)-[:HAS_FILE_SPEC]->(f)
			`, map[string]any{
				"artifactID":         string(artifact.ID()),
				"kind":               string(file.Kind()),
				"purpose":            file.Purpose(),
				"specification":      file.Specification(),
				"constraints":        file.Constraints(),
				"acceptanceCriteria": file.AcceptanceCriteria(),
			})
		}

		return run(ctx, tx, `
			MATCH (a:Artifact {id: $artifactID})
			OPTIONAL MATCH (a)-[r:HAS_FILE_SPEC]->(f:ArtifactFile)
			DELETE r, f
		`, map[string]any{
			"artifactID": string(artifact.ID()),
		})
	})
}

func (s *planService) SaveTask(ctx context.Context, task plan.Task) error {
	if task == nil {
		return fmt.Errorf("task is required")
	}

	if task.Section() == nil {
		return fmt.Errorf("task section is required")
	}

	exists, err := s.hasSection(ctx, task.Section().ID())
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("task references unknown section %q", task.Section().ID())
	}

	if task.HasParentID() {
		exists, err := s.hasTask(ctx, task.ParentID())
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("task references unknown parent task %q", task.ParentID())
		}
	}

	for _, dependencyID := range task.DependencyIDs() {
		exists, err := s.hasTask(ctx, dependencyID)
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("task references unknown dependency task %q", dependencyID)
		}
	}

	for _, subTaskID := range task.SubTaskIDs() {
		exists, err := s.hasTask(ctx, subTaskID)
		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("task references unknown sub task %q", subTaskID)
		}
	}

	return s.write(ctx, func(ctx context.Context, tx neo4j.ManagedTransaction) error {
		if err := run(ctx, tx, `
			MATCH (s:Section {id: $sectionID})
			MERGE (t:Task {id: $id})
			SET
				t.kind = $kind,
				t.action = $action,
				t.name = $name,
				t.description = $description,
				t.parentID = $parentID
			MERGE (s)-[:HAS_TASK]->(t)
		`, map[string]any{
			"id":          string(task.ID()),
			"sectionID":   string(task.Section().ID()),
			"kind":        string(task.Kind()),
			"action":      string(task.Action()),
			"name":        task.Name(),
			"description": task.Description(),
			"parentID":    string(task.ParentID()),
		}); err != nil {
			return err
		}

		if err := run(ctx, tx, `
			MATCH (t:Task {id: $id})
			OPTIONAL MATCH (t)-[dependency:DEPENDS_ON]->(:Task)
			DELETE dependency
		`, map[string]any{
			"id": string(task.ID()),
		}); err != nil {
			return err
		}

		if err := run(ctx, tx, `
			MATCH (t:Task {id: $id})
			OPTIONAL MATCH (t)-[subtask:HAS_SUBTASK]->(:Task)
			DELETE subtask
		`, map[string]any{
			"id": string(task.ID()),
		}); err != nil {
			return err
		}

		if err := run(ctx, tx, `
			MATCH (t:Task {id: $id})
			OPTIONAL MATCH (t)-[specRel:HAS_SPEC]->(spec:TaskSpec)
			DELETE specRel, spec
		`, map[string]any{
			"id": string(task.ID()),
		}); err != nil {
			return err
		}

		for _, dependencyID := range task.DependencyIDs() {
			if err := run(ctx, tx, `
				MATCH
					(t:Task {id: $taskID}),
					(dependency:Task {id: $dependencyID})
				MERGE (t)-[:DEPENDS_ON]->(dependency)
			`, map[string]any{
				"taskID":       string(task.ID()),
				"dependencyID": string(dependencyID),
			}); err != nil {
				return err
			}
		}

		for _, subTaskID := range task.SubTaskIDs() {
			if err := run(ctx, tx, `
				MATCH
					(t:Task {id: $taskID}),
					(subTask:Task {id: $subTaskID})
				MERGE (t)-[:HAS_SUBTASK]->(subTask)
			`, map[string]any{
				"taskID":    string(task.ID()),
				"subTaskID": string(subTaskID),
			}); err != nil {
				return err
			}
		}

		if task.HasSpec() && task.Spec() != nil {
			spec := task.Spec()

			return run(ctx, tx, `
				MATCH (t:Task {id: $taskID})
				MERGE (spec:TaskSpec {taskID: $taskID})
				SET
					spec.goal = $goal,
					spec.instructions = $instructions,
					spec.constraints = $constraints,
					spec.acceptanceCriteria = $acceptanceCriteria
				MERGE (t)-[:HAS_SPEC]->(spec)
			`, map[string]any{
				"taskID":             string(task.ID()),
				"goal":               spec.Goal(),
				"instructions":       spec.Instructions(),
				"constraints":        spec.Constraints(),
				"acceptanceCriteria": spec.AcceptanceCriteria(),
			})
		}

		return nil
	})
}

func (s *planService) hasProject(ctx context.Context, id plan.ProjectID) (bool, error) {
	return s.exists(ctx, "Project", string(id))
}

func (s *planService) hasVersion(ctx context.Context, id plan.VersionID) (bool, error) {
	return s.exists(ctx, "Version", string(id))
}

func (s *planService) hasSection(ctx context.Context, id plan.SectionID) (bool, error) {
	return s.exists(ctx, "Section", string(id))
}

func (s *planService) hasTask(ctx context.Context, id plan.TaskID) (bool, error) {
	return s.exists(ctx, "Task", string(id))
}

func (s *planService) exists(ctx context.Context, label string, id string) (bool, error) {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: s.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, fmt.Sprintf(`
			MATCH (node:%s {id: $id})
			RETURN count(node) > 0 AS exists
		`, label), map[string]any{
			"id": id,
		})
		if err != nil {
			return false, err
		}

		record, err := result.Single(ctx)
		if err != nil {
			return false, err
		}

		exists, ok := record.Get("exists")
		if !ok {
			return false, fmt.Errorf("exists value was not returned")
		}

		typed, ok := exists.(bool)
		if !ok {
			return false, fmt.Errorf("exists value has invalid type %T", exists)
		}

		return typed, nil
	})
	if err != nil {
		return false, err
	}

	typed, ok := value.(bool)
	if !ok {
		return false, fmt.Errorf("exists result has invalid type %T", value)
	}

	return typed, nil
}

func (s *planService) write(
	ctx context.Context,
	fn func(ctx context.Context, tx neo4j.ManagedTransaction) error,
) error {
	session := s.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: s.database,
	})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		if err := fn(ctx, tx); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}

func run(
	ctx context.Context,
	tx neo4j.ManagedTransaction,
	query string,
	params map[string]any,
) error {
	result, err := tx.Run(ctx, query, params)
	if err != nil {
		return err
	}

	_, err = result.Consume(ctx)

	return err
}
