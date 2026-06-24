package memgraph

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

type planTaskRepository struct {
	driver            neo4j.DriverWithContext
	database          string
	sectionRepository plan.SectionRepository
}

func (r *planTaskRepository) IDsByProject(
	ctx context.Context,
	projectID plan.ProjectID,
) ([]plan.TaskID, error) {
	return r.idsBy(ctx, `
		MATCH (p:Project {id: $projectID})-[:HAS_VERSION]->(:Version)-[:HAS_SECTION]->(:Section)-[:HAS_TASK]->(t:Task)
		RETURN t.id AS id
		ORDER BY t.name
	`, map[string]any{
		"projectID": string(projectID),
	})
}

func (r *planTaskRepository) IDsByVersion(
	ctx context.Context,
	versionID plan.VersionID,
) ([]plan.TaskID, error) {
	return r.idsBy(ctx, `
		MATCH (:Version {id: $versionID})-[:HAS_SECTION]->(:Section)-[:HAS_TASK]->(t:Task)
		RETURN t.id AS id
		ORDER BY t.name
	`, map[string]any{
		"versionID": string(versionID),
	})
}

func (r *planTaskRepository) Get(
	ctx context.Context,
	id plan.TaskID,
) (plan.Task, error) {
	return r.getBy(ctx, `
		MATCH (s:Section)-[:HAS_TASK]->(t:Task {id: $id})
		OPTIONAL MATCH (t)-[:DEPENDS_ON]->(dependency:Task)
		OPTIONAL MATCH (t)-[:HAS_SUBTASK]->(subTask:Task)
		OPTIONAL MATCH (t)-[:HAS_SPEC]->(spec:TaskSpec)
		RETURN
			s.id AS sectionID,
			t.kind AS kind,
			t.action AS action,
			t.name AS name,
			t.description AS description,
			t.parentID AS parentID,
			collect(DISTINCT dependency.id) AS dependencyIDs,
			collect(DISTINCT subTask.id) AS subTaskIDs,
			spec.goal AS specGoal,
			spec.instructions AS specInstructions,
			spec.constraints AS specConstraints,
			spec.acceptanceCriteria AS specAcceptanceCriteria
	`, map[string]any{
		"id": string(id),
	})
}

func (r *planTaskRepository) RootsByProject(
	ctx context.Context,
	projectID plan.ProjectID,
) ([]plan.TaskID, error) {
	return r.idsBy(ctx, `
		MATCH (p:Project {id: $projectID})-[:HAS_VERSION]->(:Version)-[:HAS_SECTION]->(:Section)-[:HAS_TASK]->(t:Task)
		WHERE t.parentID IS NULL OR t.parentID = ""
		RETURN t.id AS id
		ORDER BY t.name
	`, map[string]any{
		"projectID": string(projectID),
	})
}

func (r *planTaskRepository) RootsByVersion(
	ctx context.Context,
	versionID plan.VersionID,
) ([]plan.TaskID, error) {
	return r.idsBy(ctx, `
		MATCH (:Version {id: $versionID})-[:HAS_SECTION]->(:Section)-[:HAS_TASK]->(t:Task)
		WHERE t.parentID IS NULL OR t.parentID = ""
		RETURN t.id AS id
		ORDER BY t.name
	`, map[string]any{
		"versionID": string(versionID),
	})
}

func (r *planTaskRepository) Children(
	ctx context.Context,
	parentID plan.TaskID,
) ([]plan.TaskID, error) {
	return r.idsBy(ctx, `
		MATCH (:Task {id: $parentID})-[:HAS_SUBTASK]->(child:Task)
		RETURN child.id AS id
		ORDER BY child.name
	`, map[string]any{
		"parentID": string(parentID),
	})
}

func (r *planTaskRepository) Dependents(
	ctx context.Context,
	taskID plan.TaskID,
) ([]plan.TaskID, error) {
	return r.Downstream(ctx, taskID)
}

func (r *planTaskRepository) Upstream(
	ctx context.Context,
	taskID plan.TaskID,
) ([]plan.TaskID, error) {
	return r.idsBy(ctx, `
		MATCH (:Task {id: $taskID})-[:DEPENDS_ON]->(dependency:Task)
		RETURN dependency.id AS id
		ORDER BY dependency.name
	`, map[string]any{
		"taskID": string(taskID),
	})
}

func (r *planTaskRepository) Downstream(
	ctx context.Context,
	taskID plan.TaskID,
) ([]plan.TaskID, error) {
	return r.idsBy(ctx, `
		MATCH (dependent:Task)-[:DEPENDS_ON]->(:Task {id: $taskID})
		RETURN dependent.id AS id
		ORDER BY dependent.name
	`, map[string]any{
		"taskID": string(taskID),
	})
}

func (r *planTaskRepository) ReadyByProject(
	ctx context.Context,
	projectID plan.ProjectID,
) ([]plan.TaskID, error) {
	return r.idsBy(ctx, `
		MATCH (p:Project {id: $projectID})-[:HAS_VERSION]->(:Version)-[:HAS_SECTION]->(:Section)-[:HAS_TASK]->(t:Task)
		WHERE NOT (t)-[:DEPENDS_ON]->(:Task)
		RETURN t.id AS id
		ORDER BY t.name
	`, map[string]any{
		"projectID": string(projectID),
	})
}

func (r *planTaskRepository) ReadyByVersion(
	ctx context.Context,
	versionID plan.VersionID,
) ([]plan.TaskID, error) {
	return r.idsBy(ctx, `
		MATCH (:Version {id: $versionID})-[:HAS_SECTION]->(:Section)-[:HAS_TASK]->(t:Task)
		WHERE NOT (t)-[:DEPENDS_ON]->(:Task)
		RETURN t.id AS id
		ORDER BY t.name
	`, map[string]any{
		"versionID": string(versionID),
	})
}

func (r *planTaskRepository) Has(
	ctx context.Context,
	id plan.TaskID,
) (bool, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (t:Task {id: $id})
			RETURN count(t) > 0 AS exists
		`, map[string]any{
			"id": string(id),
		})
		if err != nil {
			return false, err
		}

		record, err := result.Single(ctx)
		if err != nil {
			return false, err
		}

		rawExists, ok := record.Get("exists")
		if !ok {
			return false, fmt.Errorf("exists value was not returned")
		}

		exists, ok := rawExists.(bool)
		if !ok {
			return false, fmt.Errorf("exists value has invalid type %T", rawExists)
		}

		return exists, nil
	})
	if err != nil {
		return false, err
	}

	exists, ok := value.(bool)
	if !ok {
		return false, fmt.Errorf("exists result has invalid type %T", value)
	}

	return exists, nil
}

func (r *planTaskRepository) idsBy(
	ctx context.Context,
	query string,
	params map[string]any,
) ([]plan.TaskID, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		ids := make([]plan.TaskID, 0)

		for result.Next(ctx) {
			rawID, ok := result.Record().Get("id")
			if !ok {
				return nil, fmt.Errorf("task id was not returned")
			}

			id, ok := rawID.(string)
			if !ok {
				return nil, fmt.Errorf("task id has invalid type %T", rawID)
			}

			ids = append(ids, plan.TaskID(id))
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return ids, nil
	})
	if err != nil {
		return nil, err
	}

	ids, ok := value.([]plan.TaskID)
	if !ok {
		return nil, fmt.Errorf("task ids result has invalid type %T", value)
	}

	return ids, nil
}

func (r *planTaskRepository) getBy(
	ctx context.Context,
	query string,
	params map[string]any,
) (plan.Task, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		records, err := result.Collect(ctx)
		if err != nil {
			return nil, err
		}

		if len(records) == 0 {
			return nil, nil
		}

		record := records[0]

		rawSectionID, ok := record.Get("sectionID")
		if !ok {
			return nil, fmt.Errorf("task section id was not returned")
		}

		sectionID, ok := rawSectionID.(string)
		if !ok {
			return nil, fmt.Errorf("task section id has invalid type %T", rawSectionID)
		}

		rawKind, ok := record.Get("kind")
		if !ok {
			return nil, fmt.Errorf("task kind was not returned")
		}

		kind, ok := rawKind.(string)
		if !ok {
			return nil, fmt.Errorf("task kind has invalid type %T", rawKind)
		}

		rawAction, ok := record.Get("action")
		if !ok {
			return nil, fmt.Errorf("task action was not returned")
		}

		action, ok := rawAction.(string)
		if !ok {
			return nil, fmt.Errorf("task action has invalid type %T", rawAction)
		}

		rawName, ok := record.Get("name")
		if !ok {
			return nil, fmt.Errorf("task name was not returned")
		}

		name, ok := rawName.(string)
		if !ok {
			return nil, fmt.Errorf("task name has invalid type %T", rawName)
		}

		rawDescription, ok := record.Get("description")
		if !ok {
			return nil, fmt.Errorf("task description was not returned")
		}

		description, ok := rawDescription.(string)
		if !ok {
			return nil, fmt.Errorf("task description has invalid type %T", rawDescription)
		}

		parentID := plan.TaskID("")
		rawParentID, ok := record.Get("parentID")
		if ok && rawParentID != nil {
			typedParentID, ok := rawParentID.(string)
			if !ok {
				return nil, fmt.Errorf("task parent id has invalid type %T", rawParentID)
			}

			parentID = plan.TaskID(typedParentID)
		}

		dependencyIDs, err := taskIDsFromRecord(record, "dependencyIDs", "task dependency ids")
		if err != nil {
			return nil, err
		}

		subTaskIDs, err := taskIDsFromRecord(record, "subTaskIDs", "task sub task ids")
		if err != nil {
			return nil, err
		}

		section, err := r.sectionRepository.Get(ctx, plan.SectionID(sectionID))
		if err != nil {
			return nil, err
		}

		if section == nil {
			return nil, fmt.Errorf("task references unknown section %q", sectionID)
		}

		builder := plan.NewTaskBuilder().
			Section(section).
			Kind(plan.TaskKind(kind)).
			Action(plan.TaskAction(action)).
			Name(name).
			Description(description)

		if parentID != "" {
			builder.ParentID(parentID)
		}

		for _, dependencyID := range dependencyIDs {
			builder.AddDependency(dependencyID)
		}

		for _, subTaskID := range subTaskIDs {
			builder.AddSubTask(subTaskID)
		}

		spec, err := taskSpecFromRecord(record)
		if err != nil {
			return nil, err
		}

		if spec != nil {
			builder.Spec(spec)
		}

		task, err := builder.Build()
		if err != nil {
			return nil, err
		}

		return task, nil
	})
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	task, ok := value.(plan.Task)
	if !ok {
		return nil, fmt.Errorf("task result has invalid type %T", value)
	}

	return task, nil
}

func taskSpecFromRecord(record *neo4j.Record) (plan.TaskSpec, error) {
	rawGoal, ok := record.Get("specGoal")
	if !ok || rawGoal == nil {
		return nil, nil
	}

	goal, ok := rawGoal.(string)
	if !ok {
		return nil, fmt.Errorf("task spec goal has invalid type %T", rawGoal)
	}

	rawInstructions, ok := record.Get("specInstructions")
	if !ok {
		return nil, fmt.Errorf("task spec instructions were not returned")
	}

	instructions, ok := rawInstructions.(string)
	if !ok {
		return nil, fmt.Errorf("task spec instructions have invalid type %T", rawInstructions)
	}

	constraints, err := stringsFromRecord(record, "specConstraints", "task spec constraints")
	if err != nil {
		return nil, err
	}

	acceptanceCriteria, err := stringsFromRecord(record, "specAcceptanceCriteria", "task spec acceptance criteria")
	if err != nil {
		return nil, err
	}

	builder := plan.NewTaskSpecBuilder().
		Goal(goal).
		Instructions(instructions)

	for _, constraint := range constraints {
		builder.AddConstraint(constraint)
	}

	for _, criteria := range acceptanceCriteria {
		builder.AddAcceptanceCriteria(criteria)
	}

	return builder.Build()
}

func taskIDsFromRecord(
	record *neo4j.Record,
	field string,
	label string,
) ([]plan.TaskID, error) {
	rawValue, ok := record.Get(field)
	if !ok || rawValue == nil {
		return []plan.TaskID{}, nil
	}

	values, ok := rawValue.([]any)
	if !ok {
		typedValues, ok := rawValue.([]string)
		if !ok {
			return nil, fmt.Errorf("%s has invalid type %T", label, rawValue)
		}

		result := make([]plan.TaskID, 0, len(typedValues))
		for _, value := range typedValues {
			if value == "" {
				continue
			}

			result = append(result, plan.TaskID(value))
		}

		return result, nil
	}

	result := make([]plan.TaskID, 0, len(values))

	for _, value := range values {
		if value == nil {
			continue
		}

		typedValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("%s contains invalid value type %T", label, value)
		}

		if typedValue == "" {
			continue
		}

		result = append(result, plan.TaskID(typedValue))
	}

	return result, nil
}
