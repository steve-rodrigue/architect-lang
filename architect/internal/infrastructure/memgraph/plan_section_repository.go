package memgraph

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

type planSectionRepository struct {
	driver            neo4j.DriverWithContext
	database          string
	versionRepository plan.VersionRepository
}

func (r *planSectionRepository) IDsByVersion(ctx context.Context, versionID plan.VersionID) ([]plan.SectionID, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (v:Version {id: $versionID})-[:HAS_SECTION]->(s:Section)
			RETURN s.id AS id
			ORDER BY s.path
		`, map[string]any{
			"versionID": string(versionID),
		})
		if err != nil {
			return nil, err
		}

		ids := make([]plan.SectionID, 0)

		for result.Next(ctx) {
			rawID, ok := result.Record().Get("id")
			if !ok {
				return nil, fmt.Errorf("section id was not returned")
			}

			id, ok := rawID.(string)
			if !ok {
				return nil, fmt.Errorf("section id has invalid type %T", rawID)
			}

			ids = append(ids, plan.SectionID(id))
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return ids, nil
	})
	if err != nil {
		return nil, err
	}

	ids, ok := value.([]plan.SectionID)
	if !ok {
		return nil, fmt.Errorf("section ids result has invalid type %T", value)
	}

	return ids, nil
}

func (r *planSectionRepository) Get(ctx context.Context, id plan.SectionID) (plan.Section, error) {
	return r.getBy(
		ctx, `
		MATCH (v:Version)-[:HAS_SECTION]->(s:Section {id: $id})
		RETURN
			v.id AS versionID,
			s.kind AS kind,
			s.name AS name,
			s.parentID AS parentID
	`, map[string]any{
			"id": string(id),
		})
}

func (r *planSectionRepository) Children(ctx context.Context, parentID plan.SectionID) ([]plan.SectionID, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (:Section {id: $parentID})-[:PARENT_OF]->(child:Section)
			RETURN child.id AS id
			ORDER BY child.path
		`, map[string]any{
			"parentID": string(parentID),
		})
		if err != nil {
			return nil, err
		}

		ids := make([]plan.SectionID, 0)

		for result.Next(ctx) {
			rawID, ok := result.Record().Get("id")
			if !ok {
				return nil, fmt.Errorf("child section id was not returned")
			}

			id, ok := rawID.(string)
			if !ok {
				return nil, fmt.Errorf("child section id has invalid type %T", rawID)
			}

			ids = append(ids, plan.SectionID(id))
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return ids, nil
	})
	if err != nil {
		return nil, err
	}

	ids, ok := value.([]plan.SectionID)
	if !ok {
		return nil, fmt.Errorf("child section ids result has invalid type %T", value)
	}

	return ids, nil
}

func (r *planSectionRepository) Has(ctx context.Context, id plan.SectionID) (bool, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (s:Section {id: $id})
			RETURN count(s) > 0 AS exists
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

func (r *planSectionRepository) getBy(
	ctx context.Context,
	query string,
	params map[string]any,
) (plan.Section, error) {
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
			return nil, fmt.Errorf("section version id was not returned")
		}

		versionID, ok := rawVersionID.(string)
		if !ok {
			return nil, fmt.Errorf("section version id has invalid type %T", rawVersionID)
		}

		rawKind, ok := record.Get("kind")
		if !ok {
			return nil, fmt.Errorf("section kind was not returned")
		}

		kind, ok := rawKind.(string)
		if !ok {
			return nil, fmt.Errorf("section kind has invalid type %T", rawKind)
		}

		rawName, ok := record.Get("name")
		if !ok {
			return nil, fmt.Errorf("section name was not returned")
		}

		name, ok := rawName.(string)
		if !ok {
			return nil, fmt.Errorf("section name has invalid type %T", rawName)
		}

		parentID := plan.SectionID("")

		rawParentID, ok := record.Get("parentID")
		if ok && rawParentID != nil {
			typedParentID, ok := rawParentID.(string)
			if !ok {
				return nil, fmt.Errorf("section parent id has invalid type %T", rawParentID)
			}

			parentID = plan.SectionID(typedParentID)
		}

		version, err := r.versionRepository.Get(ctx, plan.VersionID(versionID))
		if err != nil {
			return nil, err
		}

		if version == nil {
			return nil, fmt.Errorf("section references unknown version %q", versionID)
		}

		builder := plan.NewSectionBuilder().
			Version(version).
			Kind(plan.SectionKind(kind)).
			Name(name)

		if parentID != "" {
			builder.Parent(parentID)
		}

		section, err := builder.Build()
		if err != nil {
			return nil, err
		}

		return section, nil
	})
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	section, ok := value.(plan.Section)
	if !ok {
		return nil, fmt.Errorf("section result has invalid type %T", value)
	}

	return section, nil
}
