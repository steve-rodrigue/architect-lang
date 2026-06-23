package plan

type artifact struct {
	id      ArtifactID
	kind    ArtifactKind
	section Section
	path    string
	file    ArtifactFile
}

func (a *artifact) ID() ArtifactID {
	return a.id
}

func (a *artifact) Kind() ArtifactKind {
	return a.kind
}

func (a *artifact) Section() Section {
	return a.section
}

func (a *artifact) Path() string {
	return a.path
}

func (a *artifact) File() ArtifactFile {
	return a.file
}

func (a *artifact) IsFile() bool {
	return a.file != nil
}
