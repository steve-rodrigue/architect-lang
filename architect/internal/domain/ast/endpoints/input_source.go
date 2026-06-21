package endpoints

type inputSource struct {
	kind     InputSourceKind
	optional bool
}

func (s *inputSource) Kind() InputSourceKind {
	return s.kind
}

func (s *inputSource) IsOptional() bool {
	return s.optional
}
