package deployments

type deployment struct {
	name      string
	vendor    Vendor
	inventory string
	services  []ServiceDeployment
}

func (d *deployment) Name() string {
	return d.name
}

func (d *deployment) Vendor() Vendor {
	return d.vendor
}

func (d *deployment) Inventory() string {
	return d.inventory
}

func (d *deployment) Services() []ServiceDeployment {
	services := make([]ServiceDeployment, len(d.services))
	copy(services, d.services)

	return services
}
