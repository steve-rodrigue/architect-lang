package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"

type deployment struct {
	name      string
	vendor    deployments.Vendor
	inventory string
	services  []ServiceDeployment
	ast       deployments.Deployment
}

func (d *deployment) Name() string {
	return d.name
}

func (d *deployment) Vendor() deployments.Vendor {
	return d.vendor
}

func (d *deployment) Inventory() string {
	return d.inventory
}

func (d *deployment) Services() []ServiceDeployment {
	return append([]ServiceDeployment(nil), d.services...)
}

func (d *deployment) AST() deployments.Deployment {
	return d.ast
}
