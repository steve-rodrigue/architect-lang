package model

type version struct {
	number       string
	objects      []Object
	services     []Service
	deployments  []Deployment
	nextVersions []Version

	objectsByName     map[string]Object
	servicesByName    map[string]Service
	deploymentsByName map[string]Deployment
	eventsByName      map[string]Event
}

func (v *version) Number() string {
	return v.number
}

func (v *version) Objects() []Object {
	return append([]Object(nil), v.objects...)
}

func (v *version) Services() []Service {
	return append([]Service(nil), v.services...)
}

func (v *version) Deployments() []Deployment {
	return append([]Deployment(nil), v.deployments...)
}

func (v *version) NextVersions() []Version {
	return append([]Version(nil), v.nextVersions...)
}

func (v *version) ObjectByName(name string) Object {
	return v.objectsByName[name]
}

func (v *version) ServiceByName(name string) Service {
	return v.servicesByName[name]
}

func (v *version) DeploymentByName(name string) Deployment {
	return v.deploymentsByName[name]
}

func (v *version) EventByName(name string) Event {
	return v.eventsByName[name]
}
