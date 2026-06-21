package projects

type project struct {
	name     string
	versions []Version
}

func (p *project) Name() string {
	return p.name
}

func (p *project) Versions() []Version {
	versions := make([]Version, len(p.versions))
	copy(versions, p.versions)

	return versions
}

type projectBuilder struct {
	name     string
	versions []Version
}
