package deployments

type Vendor string

const (
	VendorDockerCompose Vendor = "docker_compose"
	VendorAnsible       Vendor = "ansible"
	VendorKubernetes    Vendor = "kubernetes"
)

// NewDeploymentBuilder creates a new deployment builder
func NewDeploymentBuilder() DeploymentBuilder {
	return &deploymentBuilder{
		services: make([]ServiceDeployment, 0),
	}
}

// NewServiceDeploymentBuilder creates a new service deployment builder
func NewServiceDeploymentBuilder() ServiceDeploymentBuilder {
	return &serviceDeploymentBuilder{}
}

// DeploymentBuilder represents a deployment builder
type DeploymentBuilder interface {
	Name(name string) DeploymentBuilder
	Vendor(vendor Vendor) DeploymentBuilder
	Inventory(file string) DeploymentBuilder
	AddService(service ServiceDeployment) DeploymentBuilder
	Build() (Deployment, error)
}

// Deployment represents a deployment
type Deployment interface {
	Name() string
	Vendor() Vendor
	Inventory() string
	Services() []ServiceDeployment
}

// ServiceDeploymentBuilder represents a service deployment builder
type ServiceDeploymentBuilder interface {
	Name(name string) ServiceDeploymentBuilder
	Replicas(count int) ServiceDeploymentBuilder
	Domain(domain string) ServiceDeploymentBuilder
	Host(host string) ServiceDeploymentBuilder
	Volume(volume string) ServiceDeploymentBuilder
	Backup(enabled bool) ServiceDeploymentBuilder
	Build() (ServiceDeployment, error)
}

// ServiceDeployment represents a service deployment
type ServiceDeployment interface {
	Name() string
	Replicas() int
	HasReplicas() bool
	Domain() string
	Host() string
	Volume() string
	Backup() bool
	HasBackup() bool
}
