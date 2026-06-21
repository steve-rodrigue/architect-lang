package sources

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/projects"
)

type versionBuilder struct {
	version     projects.Version
	objects     []objects.Object
	services    []Service
	deployments []deployments.Deployment
}

func (b *versionBuilder) Version(version projects.Version) VersionBuilder {
	b.version = version
	return b
}

func (b *versionBuilder) AddObject(object objects.Object) VersionBuilder {
	b.objects = append(b.objects, object)
	return b
}

func (b *versionBuilder) AddService(service Service) VersionBuilder {
	b.services = append(b.services, service)
	return b
}

func (b *versionBuilder) AddDeployment(deployment deployments.Deployment) VersionBuilder {
	b.deployments = append(b.deployments, deployment)
	return b
}

func (b *versionBuilder) Build() (Version, error) {
	if b.version == nil {
		return nil, fmt.Errorf("version source version is required")
	}

	for _, object := range b.objects {
		if object == nil {
			return nil, fmt.Errorf("version source object is required")
		}
	}

	for _, service := range b.services {
		if service == nil {
			return nil, fmt.Errorf("version source service is required")
		}
	}

	for _, deployment := range b.deployments {
		if deployment == nil {
			return nil, fmt.Errorf("version source deployment is required")
		}
	}

	return &version{
		version:     b.version,
		objects:     append([]objects.Object(nil), b.objects...),
		services:    append([]Service(nil), b.services...),
		deployments: append([]deployments.Deployment(nil), b.deployments...),
	}, nil
}
