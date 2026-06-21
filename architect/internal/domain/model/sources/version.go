package sources

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/projects"
)

type version struct {
	version     projects.Version
	objects     []objects.Object
	services    []Service
	deployments []deployments.Deployment
}

func (s *version) Version() projects.Version {
	return s.version
}

func (s *version) Objects() []objects.Object {
	return append([]objects.Object(nil), s.objects...)
}

func (s *version) Services() []Service {
	return append([]Service(nil), s.services...)
}

func (s *version) Deployments() []deployments.Deployment {
	return append([]deployments.Deployment(nil), s.deployments...)
}
