package plan

import "fmt"

type dependencyBuilder struct {
	version Version

	kind      DependencyKind
	scope     DependencyScope
	ecosystem DependencyEcosystem

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

func (b *dependencyBuilder) Scope(scope DependencyScope) DependencyBuilder {
	b.scope = scope
	return b
}

func (b *dependencyBuilder) Ecosystem(ecosystem DependencyEcosystem) DependencyBuilder {
	b.ecosystem = ecosystem
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

	if err := validateDependencyKind(b.kind); err != nil {
		return nil, err
	}

	if err := validateDependencyScope(b.scope); err != nil {
		return nil, err
	}

	if err := validateDependencyEcosystem(b.ecosystem); err != nil {
		return nil, err
	}

	if err := validateDependencyCombination(
		b.kind,
		b.scope,
		b.ecosystem,
	); err != nil {
		return nil, err
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
		scope:      b.scope,
		ecosystem:  b.ecosystem,
		name:       b.name,
		depVersion: b.depVersion,
		source:     b.source,
	}, nil
}

func validateDependencyKind(kind DependencyKind) error {
	if kind == "" {
		return fmt.Errorf("dependency kind is required")
	}

	for _, allowed := range DependencyKinds() {
		if allowed == kind {
			return nil
		}
	}

	return fmt.Errorf(
		"invalid dependency kind %q, allowed values are: %s",
		kind,
		DependencyKindsString(),
	)
}

func validateDependencyScope(scope DependencyScope) error {
	if scope == "" {
		return fmt.Errorf("dependency scope is required")
	}

	for _, allowed := range DependencyScopes() {
		if allowed == scope {
			return nil
		}
	}

	return fmt.Errorf(
		"invalid dependency scope %q, allowed values are: %s",
		scope,
		DependencyScopesString(),
	)
}

func validateDependencyEcosystem(ecosystem DependencyEcosystem) error {
	if ecosystem == "" {
		return fmt.Errorf("dependency ecosystem is required")
	}

	for _, allowed := range DependencyEcosystems() {
		if allowed == ecosystem {
			return nil
		}
	}

	return fmt.Errorf(
		"invalid dependency ecosystem %q, allowed values are: %s",
		ecosystem,
		DependencyEcosystemsString(),
	)
}
func validateDependencyCombination(
	kind DependencyKind,
	_ DependencyScope,
	ecosystem DependencyEcosystem,
) error {

	switch kind {

	case DependencyKindEnvironment:
		switch ecosystem {
		case DependencyEcosystemGo,
			DependencyEcosystemPython,
			DependencyEcosystemNode:
			return nil
		}

		return fmt.Errorf(
			"environment dependencies must use one of: go, python, node",
		)

	case DependencyImage:
		if ecosystem != DependencyEcosystemDocker {
			return fmt.Errorf(
				"image dependencies must use ecosystem %q",
				DependencyEcosystemDocker,
			)
		}

		return nil

	case DependencyTool:
		if ecosystem == DependencyEcosystemDocker {
			return fmt.Errorf(
				"tool dependencies cannot use ecosystem %q",
				DependencyEcosystemDocker,
			)
		}

		return nil

	case DependencyModule:
		if ecosystem == DependencyEcosystemDocker {
			return fmt.Errorf(
				"module dependencies cannot use ecosystem %q",
				DependencyEcosystemDocker,
			)
		}

		return nil
	}

	return nil
}
