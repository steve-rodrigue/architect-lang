package plan

type artifactFile struct {
	kind               ArtifactFileKind
	purpose            string
	specification      string
	constraints        []string
	acceptanceCriteria []string
}

func (f *artifactFile) Kind() ArtifactFileKind {
	return f.kind
}

func (f *artifactFile) Purpose() string {
	return f.purpose
}

func (f *artifactFile) Specification() string {
	return f.specification
}

func (f *artifactFile) Constraints() []string {
	return append([]string(nil), f.constraints...)
}

func (f *artifactFile) AcceptanceCriteria() []string {
	return append([]string(nil), f.acceptanceCriteria...)
}
