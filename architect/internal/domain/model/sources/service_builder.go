package sources

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/consumers"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"
)

type serviceBuilder struct {
	service     services.Service
	application applications.Application
	endpoints   []endpoints.Endpoint
	consumers   []consumers.Consumer
}

func (b *serviceBuilder) Service(service services.Service) ServiceBuilder {
	b.service = service
	return b
}

func (b *serviceBuilder) Application(application applications.Application) ServiceBuilder {
	b.application = application
	return b
}

func (b *serviceBuilder) AddEndpoint(endpoint endpoints.Endpoint) ServiceBuilder {
	b.endpoints = append(b.endpoints, endpoint)
	return b
}

func (b *serviceBuilder) AddConsumer(consumer consumers.Consumer) ServiceBuilder {
	b.consumers = append(b.consumers, consumer)
	return b
}

func (b *serviceBuilder) Build() (Service, error) {
	if b.service == nil {
		return nil, fmt.Errorf("service source service is required")
	}

	for _, endpoint := range b.endpoints {
		if endpoint == nil {
			return nil, fmt.Errorf("service source endpoint is required")
		}
	}

	for _, consumer := range b.consumers {
		if consumer == nil {
			return nil, fmt.Errorf("service source consumer is required")
		}
	}

	return &service{
		service:     b.service,
		application: b.application,
		endpoints:   append([]endpoints.Endpoint(nil), b.endpoints...),
		consumers:   append([]consumers.Consumer(nil), b.consumers...),
	}, nil
}
