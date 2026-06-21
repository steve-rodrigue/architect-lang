package applications

import "fmt"

type applicationBuilder struct {
	name          string
	ports         []Port
	endpointFiles []string
	consumerFiles []string
}

func (b *applicationBuilder) Name(name string) ApplicationBuilder {
	b.name = name
	return b
}

func (b *applicationBuilder) AddPort(port Port) ApplicationBuilder {
	if port != nil {
		b.ports = append(b.ports, port)
	}

	return b
}

func (b *applicationBuilder) AddEndpointFile(file string) ApplicationBuilder {
	if file != "" {
		b.endpointFiles = append(b.endpointFiles, file)
	}

	return b
}

func (b *applicationBuilder) AddConsumerFile(file string) ApplicationBuilder {
	if file != "" {
		b.consumerFiles = append(b.consumerFiles, file)
	}

	return b
}

func (b *applicationBuilder) Build() (Application, error) {
	if b.name == "" {
		return nil, fmt.Errorf("application name is required")
	}

	return &application{
		name:          b.name,
		ports:         b.ports,
		endpointFiles: b.endpointFiles,
		consumerFiles: b.consumerFiles,
	}, nil
}
