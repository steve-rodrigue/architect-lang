package applications

type PortKind string

const (
	PortKindRest   PortKind = "rest"
	PortKindEvents PortKind = "events"
)

type PortDirection string

const (
	PortDirectionEmit   PortDirection = "emit"
	PortDirectionListen PortDirection = "listen"
)

// NewApplicationBuilder creates a new application builder
func NewApplicationBuilder() ApplicationBuilder {
	return &applicationBuilder{
		ports:         make([]Port, 0),
		objectFiles:   make([]string, 0),
		endpointFiles: make([]string, 0),
		consumerFiles: make([]string, 0),
	}
}

// NewPortBuilder creates a new port builder
func NewPortBuilder() PortBuilder {
	return &portBuilder{}
}

// ApplicationBuilder represents an application builder
type ApplicationBuilder interface {
	Name(name string) ApplicationBuilder
	AddPort(port Port) ApplicationBuilder
	AddObjectFile(file string) ApplicationBuilder
	AddEndpointFile(file string) ApplicationBuilder
	AddConsumerFile(file string) ApplicationBuilder
	Build() (Application, error)
}

// Application represents an application
type Application interface {
	Name() string
	Ports() []Port
	ObjectFiles() []string
	EndpointFiles() []string
	ConsumerFiles() []string
}

// PortBuilder represents a port builder
type PortBuilder interface {
	Direction(direction PortDirection) PortBuilder
	Kind(kind PortKind) PortBuilder
	Number(number int) PortBuilder
	Build() (Port, error)
}

// Port represents a port
type Port interface {
	Direction() PortDirection
	Kind() PortKind
	Number() int
}
