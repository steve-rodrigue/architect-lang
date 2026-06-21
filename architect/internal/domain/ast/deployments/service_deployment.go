package deployments

type serviceDeployment struct {
	name        string
	replicas    int
	hasReplicas bool
	domain      string
	host        string
	volume      string
	backup      bool
	hasBackup   bool
}

func (s *serviceDeployment) Name() string {
	return s.name
}

func (s *serviceDeployment) Replicas() int {
	return s.replicas
}

func (s *serviceDeployment) HasReplicas() bool {
	return s.hasReplicas
}

func (s *serviceDeployment) Domain() string {
	return s.domain
}

func (s *serviceDeployment) Host() string {
	return s.host
}

func (s *serviceDeployment) Volume() string {
	return s.volume
}

func (s *serviceDeployment) Backup() bool {
	return s.backup
}

func (s *serviceDeployment) HasBackup() bool {
	return s.hasBackup
}
