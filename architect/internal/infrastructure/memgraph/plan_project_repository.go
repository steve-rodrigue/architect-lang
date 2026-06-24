package memgraph

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

type planProjectRepository struct {
	driver   neo4j.DriverWithContext
	database string
}

func (r *planProjectRepository) IDs(ctx context.Context) ([]plan.ProjectID, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (p:Project)
			RETURN p.id AS id
			ORDER BY p.name
		`, nil)
		if err != nil {
			return nil, err
		}

		ids := make([]plan.ProjectID, 0)

		for result.Next(ctx) {
			id, ok := result.Record().Get("id")
			if !ok {
				return nil, fmt.Errorf("project id was not returned")
			}

			typed, ok := id.(string)
			if !ok {
				return nil, fmt.Errorf("project id has invalid type %T", id)
			}

			ids = append(ids, plan.ProjectID(typed))
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return ids, nil
	})
	if err != nil {
		return nil, err
	}

	ids, ok := value.([]plan.ProjectID)
	if !ok {
		return nil, fmt.Errorf("project ids result has invalid type %T", value)
	}

	return ids, nil
}

func (r *planProjectRepository) Names(ctx context.Context) ([]string, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (p:Project)
			RETURN p.name AS name
			ORDER BY p.name
		`, nil)
		if err != nil {
			return nil, err
		}

		names := make([]string, 0)

		for result.Next(ctx) {
			name, ok := result.Record().Get("name")
			if !ok {
				return nil, fmt.Errorf("project name was not returned")
			}

			typed, ok := name.(string)
			if !ok {
				return nil, fmt.Errorf("project name has invalid type %T", name)
			}

			names = append(names, typed)
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return names, nil
	})
	if err != nil {
		return nil, err
	}

	names, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("project names result has invalid type %T", value)
	}

	return names, nil
}

func (r *planProjectRepository) Get(ctx context.Context, id plan.ProjectID) (plan.Project, error) {
	return r.getBy(
		ctx,
		`
		MATCH (p:Project {id: $id})
		RETURN
			p.name AS name,
			p.schemaVersion AS schemaVersion
	`, map[string]any{
			"id": string(id),
		})
}

func (r *planProjectRepository) GetByName(ctx context.Context, name string) (plan.Project, error) {
	return r.getBy(
		ctx,
		`
		MATCH (p:Project {name: $name})
		RETURN
			p.name AS name,
			p.schemaVersion AS schemaVersion
	`, map[string]any{
			"name": name,
		})
}

func (r *planProjectRepository) CurrentByName(ctx context.Context, name string) (plan.Project, error) {
	return r.GetByName(ctx, name)
}

func (r *planProjectRepository) Previous(ctx context.Context, id plan.ProjectID) (plan.ProjectID, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (current:Project {id: $id})
			OPTIONAL MATCH (previous:Project)-[:NEXT_PROJECT]->(current)
			RETURN previous.id AS id
		`, map[string]any{
			"id": string(id),
		})
		if err != nil {
			return nil, err
		}

		record, err := result.Single(ctx)
		if err != nil {
			if neo4j.IsNeo4jError(err) {
				return nil, err
			}

			return nil, err
		}

		rawID, ok := record.Get("id")
		if !ok || rawID == nil {
			return plan.ProjectID(""), nil
		}

		typed, ok := rawID.(string)
		if !ok {
			return nil, fmt.Errorf("previous project id has invalid type %T", rawID)
		}

		return plan.ProjectID(typed), nil
	})
	if err != nil {
		return "", err
	}

	projectID, ok := value.(plan.ProjectID)
	if !ok {
		return "", fmt.Errorf("previous project result has invalid type %T", value)
	}

	return projectID, nil
}

func (r *planProjectRepository) History(ctx context.Context, name string) ([]plan.ProjectID, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (p:Project {name: $name})
			RETURN p.id AS id
			ORDER BY p.id
		`, map[string]any{
			"name": name,
		})
		if err != nil {
			return nil, err
		}

		ids := make([]plan.ProjectID, 0)

		for result.Next(ctx) {
			rawID, ok := result.Record().Get("id")
			if !ok {
				return nil, fmt.Errorf("project history id was not returned")
			}

			typed, ok := rawID.(string)
			if !ok {
				return nil, fmt.Errorf("project history id has invalid type %T", rawID)
			}

			ids = append(ids, plan.ProjectID(typed))
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return ids, nil
	})
	if err != nil {
		return nil, err
	}

	ids, ok := value.([]plan.ProjectID)
	if !ok {
		return nil, fmt.Errorf("project history result has invalid type %T", value)
	}

	return ids, nil
}

func (r *planProjectRepository) Has(ctx context.Context, id plan.ProjectID) (bool, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (p:Project {id: $id})
			RETURN count(p) > 0 AS exists
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

	exists, ok := value.(bool)
	if !ok {
		return false, fmt.Errorf("exists result has invalid type %T", value)
	}

	return exists, nil
}

func (r *planProjectRepository) getBy(
	ctx context.Context,
	query string,
	params map[string]any,
) (plan.Project, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		record, err := result.Single(ctx)
		if err != nil {
			if neo4j.IsNeo4jError(err) {
				return nil, err
			}

			return nil, err
		}

		projectName, ok := record.Get("name")
		if !ok {
			return nil, fmt.Errorf("project name was not returned")
		}

		projectNameString, ok := projectName.(string)
		if !ok {
			return nil, fmt.Errorf("project name has invalid type %T", projectName)
		}

		schemaVersion, ok := record.Get("schemaVersion")
		if !ok {
			return nil, fmt.Errorf("project schema version was not returned")
		}

		schemaVersionString, ok := schemaVersion.(string)
		if !ok {
			return nil, fmt.Errorf("project schema version has invalid type %T", schemaVersion)
		}

		modelProject, err := model.NewProjectBuilder().
			Name(projectNameString).
			Build()
		if err != nil {
			return nil, err
		}

		project, err := plan.NewProjectBuilder().
			Model(modelProject).
			SchemaVersion(schemaVersionString).
			Build()
		if err != nil {
			return nil, err
		}

		return project, nil
	})
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	project, ok := value.(plan.Project)
	if !ok {
		return nil, fmt.Errorf("project result has invalid type %T", value)
	}

	return project, nil
}
