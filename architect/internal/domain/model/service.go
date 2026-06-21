package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"

type service struct {
	name          string
	kind          services.ServiceKind
	version       string
	image         string
	exposedPorts  []int
	application   Application
	dependencies  []Service
	storedObjects []Object
	events        []Event
	ast           services.Service
}

func (s *service) Name() string {
	return s.name
}

func (s *service) Kind() services.ServiceKind {
	return s.kind
}

func (s *service) Version() string {
	return s.version
}

func (s *service) Image() string {
	return s.image
}

func (s *service) ExposedPorts() []int {
	return append([]int(nil), s.exposedPorts...)
}

func (s *service) Application() Application {
	return s.application
}

func (s *service) HasApplication() bool {
	return s.application != nil
}

func (s *service) Dependencies() []Service {
	return append([]Service(nil), s.dependencies...)
}

func (s *service) StoredObjects() []Object {
	return append([]Object(nil), s.storedObjects...)
}

func (s *service) Events() []Event {
	return append([]Event(nil), s.events...)
}

func (s *service) AST() services.Service {
	return s.ast
}
