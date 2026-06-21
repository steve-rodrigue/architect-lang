package sources

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/consumers"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"
)

type service struct {
	service     services.Service
	application applications.Application
	endpoints   []endpoints.Endpoint
	consumers   []consumers.Consumer
}

func (s *service) Service() services.Service {
	return s.service
}

func (s *service) Application() applications.Application {
	return s.application
}

func (s *service) HasApplication() bool {
	return s.application != nil
}

func (s *service) Endpoints() []endpoints.Endpoint {
	return append([]endpoints.Endpoint(nil), s.endpoints...)
}

func (s *service) Consumers() []consumers.Consumer {
	return append([]consumers.Consumer(nil), s.consumers...)
}
