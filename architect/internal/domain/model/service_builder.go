package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"
)

type serviceBuilder struct {
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

func (b *serviceBuilder) Name(name string) ServiceBuilder {
	b.name = name
	return b
}

func (b *serviceBuilder) Kind(kind services.ServiceKind) ServiceBuilder {
	b.kind = kind
	return b
}

func (b *serviceBuilder) Version(version string) ServiceBuilder {
	b.version = version
	return b
}

func (b *serviceBuilder) Image(image string) ServiceBuilder {
	b.image = image
	return b
}

func (b *serviceBuilder) Expose(port int) ServiceBuilder {
	b.exposedPorts = append(b.exposedPorts, port)
	return b
}

func (b *serviceBuilder) Application(application Application) ServiceBuilder {
	b.application = application
	return b
}

func (b *serviceBuilder) AddDependency(service Service) ServiceBuilder {
	b.dependencies = append(b.dependencies, service)
	return b
}

func (b *serviceBuilder) AddStoredObject(object Object) ServiceBuilder {
	b.storedObjects = append(b.storedObjects, object)
	return b
}

func (b *serviceBuilder) AddEvent(event Event) ServiceBuilder {
	b.events = append(b.events, event)
	return b
}

func (b *serviceBuilder) AST(ast services.Service) ServiceBuilder {
	b.ast = ast
	return b
}

func (b *serviceBuilder) Build() (Service, error) {
	if b.name == "" {
		return nil, fmt.Errorf("service name is required")
	}

	if b.kind == "" {
		return nil, fmt.Errorf("service kind is required")
	}

	for _, port := range b.exposedPorts {
		if port <= 0 {
			return nil, fmt.Errorf("service exposed port must be greater than 0")
		}
	}

	for _, dependency := range b.dependencies {
		if dependency == nil {
			return nil, fmt.Errorf("service dependency is required")
		}
	}

	for _, object := range b.storedObjects {
		if object == nil {
			return nil, fmt.Errorf("service stored object is required")
		}
	}

	for _, event := range b.events {
		if event == nil {
			return nil, fmt.Errorf("service event is required")
		}
	}

	return &service{
		name:          b.name,
		kind:          b.kind,
		version:       b.version,
		image:         b.image,
		exposedPorts:  append([]int(nil), b.exposedPorts...),
		application:   b.application,
		dependencies:  append([]Service(nil), b.dependencies...),
		storedObjects: append([]Object(nil), b.storedObjects...),
		events:        append([]Event(nil), b.events...),
		ast:           b.ast,
	}, nil
}
