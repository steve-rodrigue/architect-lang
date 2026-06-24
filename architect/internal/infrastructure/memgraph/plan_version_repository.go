package memgraph

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

type planVersionRepository struct {
	driver            neo4j.DriverWithContext
	database          string
	projectRepository plan.ProjectRepository
}

func (r *planVersionRepository) IDsByProject(ctx context.Context, projectID plan.ProjectID) ([]plan.VersionID, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (p:Project {id: $projectID})-[:HAS_VERSION]->(v:Version)
			RETURN v.id AS id
			ORDER BY v.number
		`, map[string]any{
			"projectID": string(projectID),
		})
		if err != nil {
			return nil, err
		}

		ids := make([]plan.VersionID, 0)

		for result.Next(ctx) {
			rawID, ok := result.Record().Get("id")
			if !ok {
				return nil, fmt.Errorf("version id was not returned")
			}

			id, ok := rawID.(string)
			if !ok {
				return nil, fmt.Errorf("version id has invalid type %T", rawID)
			}

			ids = append(ids, plan.VersionID(id))
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return ids, nil
	})
	if err != nil {
		return nil, err
	}

	ids, ok := value.([]plan.VersionID)
	if !ok {
		return nil, fmt.Errorf("version ids result has invalid type %T", value)
	}

	return ids, nil
}

func (r *planVersionRepository) Get(ctx context.Context, id plan.VersionID) (plan.Version, error) {
	return r.getBy(
		ctx,
		`
		MATCH (p:Project)-[:HAS_VERSION]->(v:Version {id: $id})
		RETURN
			p.id AS projectID,
			v.number AS number
	`, map[string]any{
			"id": string(id),
		})
}

func (r *planVersionRepository) GetByNumber(ctx context.Context, projectID plan.ProjectID, number string) (plan.Version, error) {
	return r.getBy(
		ctx,
		`
		MATCH (p:Project {id: $projectID})-[:HAS_VERSION]->(v:Version {number: $number})
		RETURN
			p.id AS projectID,
			v.number AS number
	`, map[string]any{
			"projectID": string(projectID),
			"number":    number,
		})
}

func (r *planVersionRepository) Has(ctx context.Context, id plan.VersionID) (bool, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (v:Version {id: $id})
			RETURN count(v) > 0 AS exists
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

func (r *planVersionRepository) getBy(
	ctx context.Context,
	query string,
	params map[string]any,
) (plan.Version, error) {
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

		rawProjectID, ok := record.Get("projectID")
		if !ok {
			return nil, fmt.Errorf("version project id was not returned")
		}

		projectID, ok := rawProjectID.(string)
		if !ok {
			return nil, fmt.Errorf("version project id has invalid type %T", rawProjectID)
		}

		rawNumber, ok := record.Get("number")
		if !ok {
			return nil, fmt.Errorf("version number was not returned")
		}

		number, ok := rawNumber.(string)
		if !ok {
			return nil, fmt.Errorf("version number has invalid type %T", rawNumber)
		}

		project, err := r.projectRepository.Get(ctx, plan.ProjectID(projectID))
		if err != nil {
			return nil, err
		}

		if project == nil {
			return nil, fmt.Errorf("version references unknown project %q", projectID)
		}

		version, err := plan.NewVersionBuilder().
			Project(project).
			Number(number).
			Build()
		if err != nil {
			return nil, err
		}

		return version, nil
	})
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	version, ok := value.(plan.Version)
	if !ok {
		return nil, fmt.Errorf("version result has invalid type %T", value)
	}

	return version, nil
}
