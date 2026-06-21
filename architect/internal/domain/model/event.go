package model

type event struct {
	name       string
	declaredBy Service
}

func (e *event) Name() string {
	return e.name
}

func (e *event) DeclaredBy() Service {
	return e.declaredBy
}
