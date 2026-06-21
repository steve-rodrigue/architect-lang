package projects

import "fmt"

func (b *projectBuilder) Name(name string) ProjectBuilder {
	b.name = name
	return b
}

func (b *projectBuilder) AddVersion(version Version) ProjectBuilder {
	if version != nil {
		b.versions = append(b.versions, version)
	}

	return b
}

func (b *projectBuilder) Build() (Project, error) {
	if b.name == "" {
		return nil, fmt.Errorf("project name is required")
	}

	if len(b.versions) == 0 {
		return nil, fmt.Errorf("at least one project version is required")
	}

	versions := make([]Version, len(b.versions))
	copy(versions, b.versions)

	return &project{
		name:     b.name,
		versions: versions,
	}, nil
}
