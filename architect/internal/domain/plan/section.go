package plan

type section struct {
	id       SectionID
	version  Version
	kind     SectionKind
	name     string
	path     string
	parentID SectionID
}

func (s *section) ID() SectionID {
	return s.id
}

func (s *section) Version() Version {
	return s.version
}

func (s *section) Kind() SectionKind {
	return s.kind
}

func (s *section) Name() string {
	return s.name
}

func (s *section) Path() string {
	return s.path
}

func (s *section) ParentID() SectionID {
	return s.parentID
}

func (s *section) HasParentID() bool {
	return s.parentID != ""
}
