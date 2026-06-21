package sources

import (
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/consumers"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/projects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"
)

// NewProjectBuilder creates a new project source builder.
func NewProjectBuilder() ProjectBuilder {
	return &projectBuilder{
		versions: make([]Version, 0),
	}
}

// NewVersionBuilder creates a new version source builder.
func NewVersionBuilder() VersionBuilder {
	return &versionBuilder{
		objects:     make([]objects.Object, 0),
		services:    make([]Service, 0),
		deployments: make([]deployments.Deployment, 0),
	}
}

// NewServiceBuilder creates a new service source builder.
func NewServiceBuilder() ServiceBuilder {
	return &serviceBuilder{
		endpoints: make([]endpoints.Endpoint, 0),
		consumers: make([]consumers.Consumer, 0),
	}
}

// FromProject creates a project source from an AST project and resolved version sources.
func FromProject(project projects.Project, versions ...Version) (Project, error) {
	builder := NewProjectBuilder().Project(project)

	for _, version := range versions {
		builder.AddVersion(version)
	}

	return builder.Build()
}

// FromVersion creates a version source from AST values.
func FromVersion(
	version projects.Version,
	objectList []objects.Object,
	serviceList []Service,
	deploymentList []deployments.Deployment,
) (Version, error) {
	builder := NewVersionBuilder().Version(version)

	for _, object := range objectList {
		builder.AddObject(object)
	}

	for _, service := range serviceList {
		builder.AddService(service)
	}

	for _, deployment := range deploymentList {
		builder.AddDeployment(deployment)
	}

	return builder.Build()
}

// FromService creates a service source from AST values.
func FromService(
	service services.Service,
	application applications.Application,
	endpointList []endpoints.Endpoint,
	consumerList []consumers.Consumer,
) (Service, error) {
	builder := NewServiceBuilder().Service(service)

	if application != nil {
		builder.Application(application)
	}

	for _, endpoint := range endpointList {
		builder.AddEndpoint(endpoint)
	}

	for _, consumer := range consumerList {
		builder.AddConsumer(consumer)
	}

	return builder.Build()
}

// ProjectBuilder represents a project source builder.
type ProjectBuilder interface {
	Project(project projects.Project) ProjectBuilder
	AddVersion(version Version) ProjectBuilder
	Build() (Project, error)
}

type Project interface {
	Project() projects.Project
	Versions() []Version
}

// VersionBuilder represents a version source builder.
type VersionBuilder interface {
	Version(version projects.Version) VersionBuilder
	AddObject(object objects.Object) VersionBuilder
	AddService(service Service) VersionBuilder
	AddDeployment(deployment deployments.Deployment) VersionBuilder
	Build() (Version, error)
}

type Version interface {
	Version() projects.Version
	Objects() []objects.Object
	Services() []Service
	Deployments() []deployments.Deployment
}

// ServiceBuilder represents a service source builder.
type ServiceBuilder interface {
	Service(service services.Service) ServiceBuilder
	Application(application applications.Application) ServiceBuilder
	AddEndpoint(endpoint endpoints.Endpoint) ServiceBuilder
	AddConsumer(consumer consumers.Consumer) ServiceBuilder
	Build() (Service, error)
}

type Service interface {
	Service() services.Service
	Application() applications.Application
	HasApplication() bool
	Endpoints() []endpoints.Endpoint
	Consumers() []consumers.Consumer
}
