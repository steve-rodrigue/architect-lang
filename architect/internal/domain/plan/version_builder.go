package plan

import "fmt"

type versionBuilder struct {
	project Project
	number  string
}

func (b *versionBuilder) Project(project Project) VersionBuilder {
	b.project = project
	return b
}

func (b *versionBuilder) Number(number string) VersionBuilder {
	b.number = number
	return b
}

func (b *versionBuilder) Build() (Version, error) {
	if b.project == nil {
		return nil, fmt.Errorf("version project is required")
	}

	if b.number == "" {
		return nil, fmt.Errorf("version number is required")
	}

	return &version{
		id:      versionID(b.project, b.number),
		project: b.project,
		number:  b.number,
	}, nil
}
