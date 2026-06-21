package deployments

import "fmt"

type deploymentBuilder struct {
	name      string
	vendor    Vendor
	inventory string
	services  []ServiceDeployment
}

func (b *deploymentBuilder) Name(name string) DeploymentBuilder {
	b.name = name
	return b
}

func (b *deploymentBuilder) Vendor(vendor Vendor) DeploymentBuilder {
	b.vendor = vendor
	return b
}

func (b *deploymentBuilder) Inventory(file string) DeploymentBuilder {
	b.inventory = file
	return b
}

func (b *deploymentBuilder) AddService(service ServiceDeployment) DeploymentBuilder {
	if service != nil {
		b.services = append(b.services, service)
	}

	return b
}

func (b *deploymentBuilder) Build() (Deployment, error) {
	if b.name == "" {
		return nil, fmt.Errorf("deployment name is required")
	}

	if b.vendor == "" {
		return nil, fmt.Errorf("deployment vendor is required")
	}

	services := make([]ServiceDeployment, len(b.services))
	copy(services, b.services)

	return &deployment{
		name:      b.name,
		vendor:    b.vendor,
		inventory: b.inventory,
		services:  services,
	}, nil
}
