package services

type ServiceKind string

const (
	ServiceKindGo         ServiceKind = "go"
	ServiceKindPython     ServiceKind = "python"
	ServiceKindPostgres   ServiceKind = "postgres"
	ServiceKindMemgraph   ServiceKind = "memgraph"
	ServiceKindMilvus     ServiceKind = "milvus"
	ServiceKindHatchet    ServiceKind = "hatchet"
	ServiceKindNginx      ServiceKind = "nginx"
	ServiceKindRedis      ServiceKind = "redis"
	ServiceKindPrometheus ServiceKind = "prometheus"
)

// NewServiceBuilder creates a new service builder
func NewServiceBuilder() ServiceBuilder {
	return &serviceBuilder{
		exposedPorts: make([]int, 0),
		dependsOn:    make([]string, 0),
		stores:       make([]string, 0),
		events:       make([]string, 0),
	}
}

// ServiceBuilder represents a service builder
type ServiceBuilder interface {
	Name(name string) ServiceBuilder
	Kind(kind ServiceKind) ServiceBuilder
	Version(version string) ServiceBuilder
	Image(image string) ServiceBuilder
	Expose(port int) ServiceBuilder
	DependsOn(serviceName string) ServiceBuilder
	Application(file string) ServiceBuilder
	Store(objectName string) ServiceBuilder
	Event(eventName string) ServiceBuilder
	Build() (Service, error)
}

// Service represents an ecosystem service.
//
// Dependency rule:
// - If ApplicationFile() is not empty, dependencies should be inferred from that application.
// - If ApplicationFile() is empty, DependsOn() contains the explicit dependencies.
type Service interface {
	Name() string
	Kind() ServiceKind
	Version() string
	Image() string
	ExposedPorts() []int
	DependsOn() []string
	ApplicationFile() string
	Stores() []string
	Events() []string
}
