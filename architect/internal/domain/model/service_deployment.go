package model

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"

type serviceDeployment struct {
	service    Service
	replicas   int
	hasReplica bool
	domain     string
	host       string
	volume     string
	backup     bool
	hasBackup  bool
	ast        deployments.ServiceDeployment
}

func (d *serviceDeployment) Service() Service {
	return d.service
}

func (d *serviceDeployment) Replicas() int {
	return d.replicas
}

func (d *serviceDeployment) HasReplicas() bool {
	return d.hasReplica
}

func (d *serviceDeployment) Domain() string {
	return d.domain
}

func (d *serviceDeployment) Host() string {
	return d.host
}

func (d *serviceDeployment) Volume() string {
	return d.volume
}

func (d *serviceDeployment) Backup() bool {
	return d.backup
}

func (d *serviceDeployment) HasBackup() bool {
	return d.hasBackup
}

func (d *serviceDeployment) AST() deployments.ServiceDeployment {
	return d.ast
}
