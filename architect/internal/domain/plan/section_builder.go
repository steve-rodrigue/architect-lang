package plan

import "fmt"

type sectionBuilder struct {
	version  Version
	kind     SectionKind
	name     string
	parentID SectionID
}

func (b *sectionBuilder) Version(version Version) SectionBuilder {
	b.version = version
	return b
}

func (b *sectionBuilder) Kind(kind SectionKind) SectionBuilder {
	b.kind = kind
	return b
}

func (b *sectionBuilder) Name(name string) SectionBuilder {
	b.name = name
	return b
}

func (b *sectionBuilder) Parent(parent SectionID) SectionBuilder {
	b.parentID = parent
	return b
}

func (b *sectionBuilder) Build() (Section, error) {
	if b.version == nil {
		return nil, fmt.Errorf("section version is required")
	}

	if b.kind == "" {
		return nil, fmt.Errorf("section kind is required")
	}

	if b.name == "" {
		return nil, fmt.Errorf("section name is required")
	}

	return &section{
		id:       GenSectionID(b.version, b.parentID, b.kind, b.name),
		version:  b.version,
		kind:     b.kind,
		name:     b.name,
		path:     GenSectionPath(b.parentID, b.kind, b.name),
		parentID: b.parentID,
	}, nil
}
