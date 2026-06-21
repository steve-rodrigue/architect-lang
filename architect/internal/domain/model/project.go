package model

type project struct {
	name     string
	versions []Version
}

func (p *project) Name() string {
	return p.name
}

func (p *project) Versions() []Version {
	return append([]Version(nil), p.versions...)
}
