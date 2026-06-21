package deployments

import "fmt"

type serviceDeploymentBuilder struct {
	name        string
	replicas    int
	hasReplicas bool
	domain      string
	host        string
	volume      string
	backup      bool
	hasBackup   bool
}

func (b *serviceDeploymentBuilder) Name(name string) ServiceDeploymentBuilder {
	b.name = name
	return b
}

func (b *serviceDeploymentBuilder) Replicas(count int) ServiceDeploymentBuilder {
	b.replicas = count
	b.hasReplicas = true
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

func (b *serviceDeploymentBuilder) Build() (ServiceDeployment, error) {
	if b.name == "" {
		return nil, fmt.Errorf("service deployment name is required")
	}

	if b.hasReplicas && b.replicas < 0 {
		return nil, fmt.Errorf("service deployment replicas cannot be negative")
	}

	return &serviceDeployment{
		name:        b.name,
		replicas:    b.replicas,
		hasReplicas: b.hasReplicas,
		domain:      b.domain,
		host:        b.host,
		volume:      b.volume,
		backup:      b.backup,
		hasBackup:   b.hasBackup,
	}, nil
}
