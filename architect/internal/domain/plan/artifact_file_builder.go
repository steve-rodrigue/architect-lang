package plan

import "fmt"

type artifactFileBuilder struct {
	kind               ArtifactFileKind
	purpose            string
	specification      string
	constraints        []string
	acceptanceCriteria []string
}

func (b *artifactFileBuilder) Kind(kind ArtifactFileKind) ArtifactFileBuilder {
	b.kind = kind
	return b
}

func (b *artifactFileBuilder) Purpose(purpose string) ArtifactFileBuilder {
	b.purpose = purpose
	return b
}

func (b *artifactFileBuilder) Specification(specification string) ArtifactFileBuilder {
	b.specification = specification
	return b
}

func (b *artifactFileBuilder) AddConstraint(constraint string) ArtifactFileBuilder {
	b.constraints = append(b.constraints, constraint)
	return b
}

func (b *artifactFileBuilder) AddAcceptanceCriteria(criteria string) ArtifactFileBuilder {
	b.acceptanceCriteria = append(b.acceptanceCriteria, criteria)
	return b
}

func (b *artifactFileBuilder) Build() (ArtifactFile, error) {
	if b.kind == "" {
		return nil, fmt.Errorf("artifact file kind is required")
	}

	if b.purpose == "" {
		return nil, fmt.Errorf("artifact file purpose is required")
	}

	if b.specification == "" {
		return nil, fmt.Errorf("artifact file specification is required")
	}

	return &artifactFile{
		kind:               b.kind,
		purpose:            b.purpose,
		specification:      b.specification,
		constraints:        append([]string(nil), b.constraints...),
		acceptanceCriteria: append([]string(nil), b.acceptanceCriteria...),
	}, nil
}
