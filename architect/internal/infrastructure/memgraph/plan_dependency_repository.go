package memgraph

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

type planDependencyRepository struct {
	driver            neo4j.DriverWithContext
	database          string
	versionRepository plan.VersionRepository
}

func (r *planDependencyRepository) IDsByProject(ctx context.Context, projectID plan.ProjectID) ([]plan.DependencyID, error) {
	return r.idsBy(
		ctx,
		`
		MATCH (p:Project {id: $projectID})-[:HAS_VERSION]->(:Version)-[:HAS_DEPENDENCY]->(d:Dependency)
		RETURN d.id AS id
		ORDER BY d.kind, d.name, d.depVersion
	`, map[string]any{
			"projectID": string(projectID),
		})
}

func (r *planDependencyRepository) IDsByVersion(ctx context.Context, versionID plan.VersionID) ([]plan.DependencyID, error) {
	return r.idsBy(
		ctx,
		`
		MATCH (:Version {id: $versionID})-[:HAS_DEPENDENCY]->(d:Dependency)
		RETURN d.id AS id
		ORDER BY d.kind, d.name, d.depVersion
	`, map[string]any{
			"versionID": string(versionID),
		})
}

func (r *planDependencyRepository) Get(ctx context.Context, id plan.DependencyID) (plan.Dependency, error) {
	return r.getBy(
		ctx,
		`
		MATCH (v:Version)-[:HAS_DEPENDENCY]->(d:Dependency {id: $id})
		RETURN
			v.id AS versionID,
			d.kind AS kind,
			d.name AS name,
			d.depVersion AS depVersion,
			d.source AS source
	`, map[string]any{
			"id": string(id),
		})
}

func (r *planDependencyRepository) ByKind(ctx context.Context, kind plan.DependencyKind) ([]plan.DependencyID, error) {
	return r.idsBy(ctx,
		`
		MATCH (d:Dependency {kind: $kind})
		RETURN d.id AS id
		ORDER BY d.name, d.depVersion
	`, map[string]any{
			"kind": string(kind),
		})
}

func (r *planDependencyRepository) Has(ctx context.Context, id plan.DependencyID) (bool, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (d:Dependency {id: $id})
			RETURN count(d) > 0 AS exists
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

func (r *planDependencyRepository) idsBy(
	ctx context.Context,
	query string,
	params map[string]any,
) ([]plan.DependencyID, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		ids := make([]plan.DependencyID, 0)

		for result.Next(ctx) {
			rawID, ok := result.Record().Get("id")
			if !ok {
				return nil, fmt.Errorf("dependency id was not returned")
			}

			id, ok := rawID.(string)
			if !ok {
				return nil, fmt.Errorf("dependency id has invalid type %T", rawID)
			}

			ids = append(ids, plan.DependencyID(id))
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return ids, nil
	})
	if err != nil {
		return nil, err
	}

	ids, ok := value.([]plan.DependencyID)
	if !ok {
		return nil, fmt.Errorf("dependency ids result has invalid type %T", value)
	}

	return ids, nil
}

func (r *planDependencyRepository) getBy(
	ctx context.Context,
	query string,
	params map[string]any,
) (plan.Dependency, error) {
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

		rawVersionID, ok := record.Get("versionID")
		if !ok {
			return nil, fmt.Errorf("dependency version id was not returned")
		}

		versionID, ok := rawVersionID.(string)
		if !ok {
			return nil, fmt.Errorf("dependency version id has invalid type %T", rawVersionID)
		}

		rawKind, ok := record.Get("kind")
		if !ok {
			return nil, fmt.Errorf("dependency kind was not returned")
		}

		kind, ok := rawKind.(string)
		if !ok {
			return nil, fmt.Errorf("dependency kind has invalid type %T", rawKind)
		}

		rawName, ok := record.Get("name")
		if !ok {
			return nil, fmt.Errorf("dependency name was not returned")
		}

		name, ok := rawName.(string)
		if !ok {
			return nil, fmt.Errorf("dependency name has invalid type %T", rawName)
		}

		rawDepVersion, ok := record.Get("depVersion")
		if !ok {
			return nil, fmt.Errorf("dependency dep version was not returned")
		}

		depVersion, ok := rawDepVersion.(string)
		if !ok {
			return nil, fmt.Errorf("dependency dep version has invalid type %T", rawDepVersion)
		}

		source := ""
		rawSource, ok := record.Get("source")
		if ok && rawSource != nil {
			typedSource, ok := rawSource.(string)
			if !ok {
				return nil, fmt.Errorf("dependency source has invalid type %T", rawSource)
			}

			source = typedSource
		}

		version, err := r.versionRepository.Get(ctx, plan.VersionID(versionID))
		if err != nil {
			return nil, err
		}

		if version == nil {
			return nil, fmt.Errorf("dependency references unknown version %q", versionID)
		}

		dependency, err := plan.NewDependencyBuilder().
			Version(version).
			Kind(plan.DependencyKind(kind)).
			Name(name).
			DepVersion(depVersion).
			Source(source).
			Build()
		if err != nil {
			return nil, err
		}

		return dependency, nil
	})
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	dependency, ok := value.(plan.Dependency)
	if !ok {
		return nil, fmt.Errorf("dependency result has invalid type %T", value)
	}

	return dependency, nil
}
