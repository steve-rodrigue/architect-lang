package services

type service struct {
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

func (s *service) Name() string {
	return s.name
}

func (s *service) Kind() ServiceKind {
	return s.kind
}

func (s *service) Version() string {
	return s.version
}

func (s *service) Image() string {
	return s.image
}

func (s *service) ExposedPorts() []int {
	ports := make([]int, len(s.exposedPorts))
	copy(ports, s.exposedPorts)

	return ports
}

func (s *service) DependsOn() []string {
	dependencies := make([]string, len(s.dependsOn))
	copy(dependencies, s.dependsOn)

	return dependencies
}

func (s *service) ApplicationFile() string {
	return s.applicationFile
}

func (s *service) Stores() []string {
	stores := make([]string, len(s.stores))
	copy(stores, s.stores)

	return stores
}

func (s *service) Events() []string {
	events := make([]string, len(s.events))
	copy(events, s.events)

	return events
}
