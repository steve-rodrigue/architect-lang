package applications

type application struct {
	name          string
	ports         []Port
	endpointFiles []string
	consumerFiles []string
}

func (a *application) Name() string {
	return a.name
}

func (a *application) Ports() []Port {
	return a.ports
}

func (a *application) EndpointFiles() []string {
	return a.endpointFiles
}

func (a *application) ConsumerFiles() []string {
	return a.consumerFiles
}
