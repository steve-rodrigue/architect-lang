package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
)

type applicationBuilder struct {
	name      string
	ports     []applications.Port
	endpoints []Endpoint
	consumers []Consumer
	ast       applications.Application
}

func (b *applicationBuilder) Name(name string) ApplicationBuilder {
	b.name = name
	return b
}

func (b *applicationBuilder) AddPort(port applications.Port) ApplicationBuilder {
	b.ports = append(b.ports, port)
	return b
}

func (b *applicationBuilder) AddEndpoint(endpoint Endpoint) ApplicationBuilder {
	b.endpoints = append(b.endpoints, endpoint)
	return b
}

func (b *applicationBuilder) AddConsumer(consumer Consumer) ApplicationBuilder {
	b.consumers = append(b.consumers, consumer)
	return b
}

func (b *applicationBuilder) AST(ast applications.Application) ApplicationBuilder {
	b.ast = ast
	return b
}

func (b *applicationBuilder) Build() (Application, error) {
	if b.name == "" {
		return nil, fmt.Errorf("application name is required")
	}

	for _, port := range b.ports {
		if port == nil {
			return nil, fmt.Errorf("application port is required")
		}
	}

	for _, endpoint := range b.endpoints {
		if endpoint == nil {
			return nil, fmt.Errorf("application endpoint is required")
		}
	}

	for _, consumer := range b.consumers {
		if consumer == nil {
			return nil, fmt.Errorf("application consumer is required")
		}
	}

	return &application{
		name:      b.name,
		ports:     append([]applications.Port(nil), b.ports...),
		endpoints: append([]Endpoint(nil), b.endpoints...),
		consumers: append([]Consumer(nil), b.consumers...),
		ast:       b.ast,
	}, nil
}
