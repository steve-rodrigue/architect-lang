package model

import "fmt"

type projectBuilder struct {
	name     string
	versions []Version
}

func (b *projectBuilder) Name(name string) ProjectBuilder {
	b.name = name
	return b
}

func (b *projectBuilder) AddVersion(version Version) ProjectBuilder {
	b.versions = append(b.versions, version)
	return b
}

func (b *projectBuilder) Build() (Project, error) {
	if b.name == "" {
		return nil, fmt.Errorf("project name is required")
	}

	versionsByNumber := make(map[string]Version)

	for _, version := range b.versions {
		if version == nil {
			return nil, fmt.Errorf("project version is required")
		}

		if _, exists := versionsByNumber[version.Number()]; exists {
			return nil, fmt.Errorf("duplicate version %q", version.Number())
		}

		versionsByNumber[version.Number()] = version
	}

	return &project{
		name:     b.name,
		versions: append([]Version(nil), b.versions...),
	}, nil
}
