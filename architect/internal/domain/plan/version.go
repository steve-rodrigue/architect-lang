package plan

type version struct {
	id      VersionID
	project Project
	number  string
}

func (v *version) ID() VersionID {
	return v.id
}

func (v *version) Project() Project {
	return v.project
}

func (v *version) Number() string {
	return v.number
}
