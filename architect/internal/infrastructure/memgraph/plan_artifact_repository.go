package memgraph

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

type planArtifactRepository struct {
	driver            neo4j.DriverWithContext
	database          string
	sectionRepository plan.SectionRepository
}

func (r *planArtifactRepository) IDsByProject(
	ctx context.Context,
	projectID plan.ProjectID,
) ([]plan.ArtifactID, error) {
	return r.idsBy(ctx, `
		MATCH (p:Project {id: $projectID})-[:HAS_VERSION]->(:Version)-[:HAS_SECTION]->(:Section)-[:HAS_ARTIFACT]->(a:Artifact)
		RETURN a.id AS id
		ORDER BY a.path
	`, map[string]any{
		"projectID": string(projectID),
	})
}

func (r *planArtifactRepository) IDsByVersion(
	ctx context.Context,
	versionID plan.VersionID,
) ([]plan.ArtifactID, error) {
	return r.idsBy(ctx, `
		MATCH (:Version {id: $versionID})-[:HAS_SECTION]->(:Section)-[:HAS_ARTIFACT]->(a:Artifact)
		RETURN a.id AS id
		ORDER BY a.path
	`, map[string]any{
		"versionID": string(versionID),
	})
}

func (r *planArtifactRepository) Get(
	ctx context.Context,
	id plan.ArtifactID,
) (plan.Artifact, error) {
	return r.getBy(ctx, `
		MATCH (s:Section)-[:HAS_ARTIFACT]->(a:Artifact {id: $id})
		OPTIONAL MATCH (a)-[:HAS_FILE_SPEC]->(f:ArtifactFile)
		RETURN
			s.id AS sectionID,
			a.kind AS kind,
			a.path AS path,
			a.isFile AS isFile,
			f.kind AS fileKind,
			f.purpose AS filePurpose,
			f.specification AS fileSpecification,
			f.constraints AS fileConstraints,
			f.acceptanceCriteria AS fileAcceptanceCriteria
	`, map[string]any{
		"id": string(id),
	})
}

func (r *planArtifactRepository) ByKind(
	ctx context.Context,
	kind plan.ArtifactKind,
) ([]plan.ArtifactID, error) {
	return r.idsBy(ctx, `
		MATCH (a:Artifact {kind: $kind})
		RETURN a.id AS id
		ORDER BY a.path
	`, map[string]any{
		"kind": string(kind),
	})
}

func (r *planArtifactRepository) Has(
	ctx context.Context,
	id plan.ArtifactID,
) (bool, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, `
			MATCH (a:Artifact {id: $id})
			RETURN count(a) > 0 AS exists
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

func (r *planArtifactRepository) idsBy(
	ctx context.Context,
	query string,
	params map[string]any,
) ([]plan.ArtifactID, error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{
		DatabaseName: r.database,
	})
	defer session.Close(ctx)

	value, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		ids := make([]plan.ArtifactID, 0)

		for result.Next(ctx) {
			rawID, ok := result.Record().Get("id")
			if !ok {
				return nil, fmt.Errorf("artifact id was not returned")
			}

			id, ok := rawID.(string)
			if !ok {
				return nil, fmt.Errorf("artifact id has invalid type %T", rawID)
			}

			ids = append(ids, plan.ArtifactID(id))
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return ids, nil
	})
	if err != nil {
		return nil, err
	}

	ids, ok := value.([]plan.ArtifactID)
	if !ok {
		return nil, fmt.Errorf("artifact ids result has invalid type %T", value)
	}

	return ids, nil
}

func (r *planArtifactRepository) getBy(
	ctx context.Context,
	query string,
	params map[string]any,
) (plan.Artifact, error) {
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
			return nil, fmt.Errorf("artifact section id was not returned")
		}

		sectionID, ok := rawSectionID.(string)
		if !ok {
			return nil, fmt.Errorf("artifact section id has invalid type %T", rawSectionID)
		}

		rawKind, ok := record.Get("kind")
		if !ok {
			return nil, fmt.Errorf("artifact kind was not returned")
		}

		kind, ok := rawKind.(string)
		if !ok {
			return nil, fmt.Errorf("artifact kind has invalid type %T", rawKind)
		}

		rawPath, ok := record.Get("path")
		if !ok {
			return nil, fmt.Errorf("artifact path was not returned")
		}

		path, ok := rawPath.(string)
		if !ok {
			return nil, fmt.Errorf("artifact path has invalid type %T", rawPath)
		}

		section, err := r.sectionRepository.Get(ctx, plan.SectionID(sectionID))
		if err != nil {
			return nil, err
		}

		if section == nil {
			return nil, fmt.Errorf("artifact references unknown section %q", sectionID)
		}

		builder := plan.NewArtifactBuilder().
			Section(section).
			Path(path)

		if plan.ArtifactKind(kind) == plan.ArtifactKindFile {
			file, err := artifactFileFromRecord(record)
			if err != nil {
				return nil, err
			}

			if file == nil {
				return nil, fmt.Errorf("artifact file %q is missing file specification", path)
			}

			builder.File(file)
		}

		artifact, err := builder.Build()
		if err != nil {
			return nil, err
		}

		return artifact, nil
	})
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	artifact, ok := value.(plan.Artifact)
	if !ok {
		return nil, fmt.Errorf("artifact result has invalid type %T", value)
	}

	return artifact, nil
}

func artifactFileFromRecord(record *neo4j.Record) (plan.ArtifactFile, error) {
	rawKind, ok := record.Get("fileKind")
	if !ok || rawKind == nil {
		return nil, nil
	}

	kind, ok := rawKind.(string)
	if !ok {
		return nil, fmt.Errorf("artifact file kind has invalid type %T", rawKind)
	}

	rawPurpose, ok := record.Get("filePurpose")
	if !ok {
		return nil, fmt.Errorf("artifact file purpose was not returned")
	}

	purpose, ok := rawPurpose.(string)
	if !ok {
		return nil, fmt.Errorf("artifact file purpose has invalid type %T", rawPurpose)
	}

	rawSpecification, ok := record.Get("fileSpecification")
	if !ok {
		return nil, fmt.Errorf("artifact file specification was not returned")
	}

	specification, ok := rawSpecification.(string)
	if !ok {
		return nil, fmt.Errorf("artifact file specification has invalid type %T", rawSpecification)
	}

	constraints, err := stringsFromRecord(record, "fileConstraints", "artifact file constraints")
	if err != nil {
		return nil, err
	}

	acceptanceCriteria, err := stringsFromRecord(record, "fileAcceptanceCriteria", "artifact file acceptance criteria")
	if err != nil {
		return nil, err
	}

	builder := plan.NewArtifactFileBuilder().
		Kind(plan.ArtifactFileKind(kind)).
		Purpose(purpose).
		Specification(specification)

	for _, constraint := range constraints {
		builder.AddConstraint(constraint)
	}

	for _, criteria := range acceptanceCriteria {
		builder.AddAcceptanceCriteria(criteria)
	}

	file, err := builder.Build()
	if err != nil {
		return nil, err
	}

	return file, nil
}

func stringsFromRecord(
	record *neo4j.Record,
	field string,
	label string,
) ([]string, error) {
	rawValue, ok := record.Get(field)
	if !ok || rawValue == nil {
		return []string{}, nil
	}

	values, ok := rawValue.([]any)
	if !ok {
		typedValues, ok := rawValue.([]string)
		if !ok {
			return nil, fmt.Errorf("%s has invalid type %T", label, rawValue)
		}

		return append([]string(nil), typedValues...), nil
	}

	result := make([]string, 0, len(values))

	for _, value := range values {
		typedValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("%s contains invalid value type %T", label, value)
		}

		result = append(result, typedValue)
	}

	return result, nil
}
