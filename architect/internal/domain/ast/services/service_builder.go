package services

import "fmt"

type serviceBuilder struct {
	name            string
	kind            ServiceKind
	version         string
	image           string
	exposedPorts    []int
	dependsOn       []string
	applicationFile string
	stores          []string
	events          []string
}

func (b *serviceBuilder) Name(name string) ServiceBuilder {
	b.name = name
	return b
}

func (b *serviceBuilder) Kind(kind ServiceKind) ServiceBuilder {
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

func (b *serviceBuilder) DependsOn(serviceName string) ServiceBuilder {
	if serviceName != "" {
		b.dependsOn = append(b.dependsOn, serviceName)
	}

	return b
}

func (b *serviceBuilder) Application(file string) ServiceBuilder {
	b.applicationFile = file
	return b
}

func (b *serviceBuilder) Store(objectName string) ServiceBuilder {
	if objectName != "" {
		b.stores = append(b.stores, objectName)
	}

	return b
}

func (b *serviceBuilder) Event(eventName string) ServiceBuilder {
	if eventName != "" {
		b.events = append(b.events, eventName)
	}

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

	if b.applicationFile != "" && len(b.dependsOn) > 0 {
		return nil, fmt.Errorf("service with application file cannot declare explicit dependencies")
	}

	exposedPorts := make([]int, len(b.exposedPorts))
	copy(exposedPorts, b.exposedPorts)

	dependsOn := make([]string, len(b.dependsOn))
	copy(dependsOn, b.dependsOn)

	stores := make([]string, len(b.stores))
	copy(stores, b.stores)

	events := make([]string, len(b.events))
	copy(events, b.events)

	return &service{
		name:            b.name,
		kind:            b.kind,
		version:         b.version,
		image:           b.image,
		exposedPorts:    exposedPorts,
		dependsOn:       dependsOn,
		applicationFile: b.applicationFile,
		stores:          stores,
		events:          events,
	}, nil
}
