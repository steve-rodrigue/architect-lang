package model

import (
	"fmt"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
)

type serviceDeploymentBuilder struct {
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

func (b *serviceDeploymentBuilder) Service(service Service) ServiceDeploymentBuilder {
	b.service = service
	return b
}

func (b *serviceDeploymentBuilder) Replicas(count int) ServiceDeploymentBuilder {
	b.replicas = count
	b.hasReplica = true
	return b
}

func (b *serviceDeploymentBuilder) Domain(domain string) ServiceDeploymentBuilder {
	b.domain = domain
	return b
}

func (b *serviceDeploymentBuilder) Host(host string) ServiceDeploymentBuilder {
	b.host = host
	return b
}

func (b *serviceDeploymentBuilder) Volume(volume string) ServiceDeploymentBuilder {
	b.volume = volume
	return b
}

func (b *serviceDeploymentBuilder) Backup(enabled bool) ServiceDeploymentBuilder {
	b.backup = enabled
	b.hasBackup = true
	return b
}

func (b *serviceDeploymentBuilder) AST(ast deployments.ServiceDeployment) ServiceDeploymentBuilder {
	b.ast = ast
	return b
}

func (b *serviceDeploymentBuilder) Build() (ServiceDeployment, error) {
	if b.service == nil {
		return nil, fmt.Errorf("service deployment service is required")
	}

	if b.hasReplica && b.replicas < 0 {
		return nil, fmt.Errorf("service deployment replicas cannot be negative")
	}

	return &serviceDeployment{
		service:    b.service,
		replicas:   b.replicas,
		hasReplica: b.hasReplica,
		domain:     b.domain,
		host:       b.host,
		volume:     b.volume,
		backup:     b.backup,
		hasBackup:  b.hasBackup,
		ast:        b.ast,
	}, nil
}
