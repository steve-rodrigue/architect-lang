package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
)

type deploymentBuilder struct {
	name      string
	vendor    deployments.Vendor
	inventory string
	services  []ServiceDeployment
	ast       deployments.Deployment
}

func (b *deploymentBuilder) Name(name string) DeploymentBuilder {
	b.name = name
	return b
}

func (b *deploymentBuilder) Vendor(vendor deployments.Vendor) DeploymentBuilder {
	b.vendor = vendor
	return b
}

func (b *deploymentBuilder) Inventory(file string) DeploymentBuilder {
	b.inventory = file
	return b
}

func (b *deploymentBuilder) AddService(service ServiceDeployment) DeploymentBuilder {
	b.services = append(b.services, service)
	return b
}

func (b *deploymentBuilder) AST(ast deployments.Deployment) DeploymentBuilder {
	b.ast = ast
	return b
}

func (b *deploymentBuilder) Build() (Deployment, error) {
	if b.name == "" {
		return nil, fmt.Errorf("deployment name is required")
	}

	if b.vendor == "" {
		return nil, fmt.Errorf("deployment vendor is required")
	}

	for _, service := range b.services {
		if service == nil {
			return nil, fmt.Errorf("deployment service is required")
		}
	}

	return &deployment{
		name:      b.name,
		vendor:    b.vendor,
		inventory: b.inventory,
		services:  append([]ServiceDeployment(nil), b.services...),
		ast:       b.ast,
	}, nil
}
