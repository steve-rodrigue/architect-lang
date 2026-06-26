package plan

type dependency struct {
	id      DependencyID
	version Version

	kind      DependencyKind
	scope     DependencyScope
	ecosystem DependencyEcosystem

	name       string
	depVersion string
	source     string
}

func (d *dependency) ID() DependencyID {
	return d.id
}

func (d *dependency) Version() Version {
	return d.version
}

func (d *dependency) Kind() DependencyKind {
	return d.kind
}

func (d *dependency) Scope() DependencyScope {
	return d.scope
}

func (d *dependency) Ecosystem() DependencyEcosystem {
	return d.ecosystem
}

func (d *dependency) Name() string {
	return d.name
}

func (d *dependency) DepVersion() string {
	return d.depVersion
}

func (d *dependency) Source() string {
	return d.source
}
