package plan

type taskSpec struct {
	goal               string
	instructions       string
	constraints        []string
	acceptanceCriteria []string
}

func (s *taskSpec) Goal() string {
	return s.goal
}

func (s *taskSpec) Instructions() string {
	return s.instructions
}

func (s *taskSpec) Constraints() []string {
	return append([]string(nil), s.constraints...)
}

func (s *taskSpec) AcceptanceCriteria() []string {
	return append([]string(nil), s.acceptanceCriteria...)
}
