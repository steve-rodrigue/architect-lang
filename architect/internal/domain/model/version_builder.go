package model

import "fmt"

type versionBuilder struct {
	number       string
	objects      []Object
	services     []Service
	deployments  []Deployment
	nextVersions []Version
}

func (b *versionBuilder) Number(number string) VersionBuilder {
	b.number = number
	return b
}

func (b *versionBuilder) AddObject(object Object) VersionBuilder {
	b.objects = append(b.objects, object)
	return b
}

func (b *versionBuilder) AddService(service Service) VersionBuilder {
	b.services = append(b.services, service)
	return b
}

func (b *versionBuilder) AddDeployment(deployment Deployment) VersionBuilder {
	b.deployments = append(b.deployments, deployment)
	return b
}

func (b *versionBuilder) AddNextVersion(version Version) VersionBuilder {
	b.nextVersions = append(b.nextVersions, version)
	return b
}

func (b *versionBuilder) Build() (Version, error) {
	if b.number == "" {
		return nil, fmt.Errorf("version number is required")
	}

	objectsByName := make(map[string]Object)
	for _, object := range b.objects {
		if object == nil {
			return nil, fmt.Errorf("version object is required")
		}

		if _, exists := objectsByName[object.Name()]; exists {
			return nil, fmt.Errorf("duplicate object %q", object.Name())
		}

		objectsByName[object.Name()] = object
	}

	servicesByName := make(map[string]Service)
	eventsByName := make(map[string]Event)
	for _, service := range b.services {
		if service == nil {
			return nil, fmt.Errorf("version service is required")
		}

		if _, exists := servicesByName[service.Name()]; exists {
			return nil, fmt.Errorf("duplicate service %q", service.Name())
		}

		servicesByName[service.Name()] = service

		for _, event := range service.Events() {
			if event == nil {
				return nil, fmt.Errorf("version service event is required")
			}

			if _, exists := eventsByName[event.Name()]; !exists {
				eventsByName[event.Name()] = event
			}
		}
	}

	deploymentsByName := make(map[string]Deployment)
	for _, deployment := range b.deployments {
		if deployment == nil {
			return nil, fmt.Errorf("version deployment is required")
		}

		if _, exists := deploymentsByName[deployment.Name()]; exists {
			return nil, fmt.Errorf("duplicate deployment %q", deployment.Name())
		}

		deploymentsByName[deployment.Name()] = deployment
	}

	for _, version := range b.nextVersions {
		if version == nil {
			return nil, fmt.Errorf("next version is required")
		}
	}

	return &version{
		number:            b.number,
		objects:           append([]Object(nil), b.objects...),
		services:          append([]Service(nil), b.services...),
		deployments:       append([]Deployment(nil), b.deployments...),
		nextVersions:      append([]Version(nil), b.nextVersions...),
		objectsByName:     objectsByName,
		servicesByName:    servicesByName,
		deploymentsByName: deploymentsByName,
		eventsByName:      eventsByName,
	}, nil
}
