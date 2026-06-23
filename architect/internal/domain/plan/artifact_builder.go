package plan

import "fmt"

type artifactBuilder struct {
	section Section
	path    string
	file    ArtifactFile
}

func (b *artifactBuilder) Section(section Section) ArtifactBuilder {
	b.section = section
	return b
}

func (b *artifactBuilder) Path(path string) ArtifactBuilder {
	b.path = path
	return b
}

func (b *artifactBuilder) File(file ArtifactFile) ArtifactBuilder {
	b.file = file
	return b
}

func (b *artifactBuilder) Build() (Artifact, error) {
	if b.section == nil {
		return nil, fmt.Errorf("artifact section is required")
	}

	if b.path == "" {
		return nil, fmt.Errorf("artifact path is required")
	}

	kind := ArtifactKindDirectory
	if b.file != nil {
		kind = ArtifactKindFile
	}

	return &artifact{
		id:      artifactID(b.section, b.path),
		kind:    kind,
		section: b.section,
		path:    cleanPath(b.path),
		file:    b.file,
	}, nil
}
