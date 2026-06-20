package objects

type value struct {
	kind ValueKind
	raw  string
}

func (v *value) Kind() ValueKind {
	return v.kind
}

func (v *value) Raw() string {
	return v.raw
}
