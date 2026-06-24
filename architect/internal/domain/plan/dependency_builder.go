package plan

import "fmt"

type dependencyBuilder struct {
	version    Version
	kind       DependencyKind
	name       string
	depVersion string
	source     string
}

func (b *dependencyBuilder) Version(version Version) DependencyBuilder {
	b.version = version
	return b
}

func (b *dependencyBuilder) Kind(kind DependencyKind) DependencyBuilder {
	b.kind = kind
	return b
}

func (b *dependencyBuilder) Name(name string) DependencyBuilder {
	b.name = name
	return b
}

func (b *dependencyBuilder) DepVersion(depVersion string) DependencyBuilder {
	b.depVersion = depVersion
	return b
}

func (b *dependencyBuilder) Source(source string) DependencyBuilder {
	b.source = source
	return b
}

func (b *dependencyBuilder) Build() (Dependency, error) {
	if b.version == nil {
		return nil, fmt.Errorf("dependency version is required")
	}

	if b.kind == "" {
		return nil, fmt.Errorf("dependency kind is required")
	}

	if b.name == "" {
		return nil, fmt.Errorf("dependency name is required")
	}

	if b.depVersion == "" {
		return nil, fmt.Errorf("dependency version is required")
	}

	return &dependency{
		id:         GenDependencyID(b.version, b.kind, b.name, b.depVersion),
		version:    b.version,
		kind:       b.kind,
		name:       b.name,
		depVersion: b.depVersion,
		source:     b.source,
	}, nil
}
