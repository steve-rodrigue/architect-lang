package projects

type version struct {
	number          string
	serviceFiles    []string
	objectFiles     []string
	deploymentFiles []string
	nextVersions    []Version
}

func (v *version) Number() string {
	return v.number
}

func (v *version) ServiceFiles() []string {
	files := make([]string, len(v.serviceFiles))
	copy(files, v.serviceFiles)

	return files
}

func (v *version) ObjectFiles() []string {
	files := make([]string, len(v.objectFiles))
	copy(files, v.objectFiles)

	return files
}

func (v *version) DeploymentFiles() []string {
	files := make([]string, len(v.deploymentFiles))
	copy(files, v.deploymentFiles)

	return files
}

func (v *version) NextVersions() []Version {
	versions := make([]Version, len(v.nextVersions))
	copy(versions, v.nextVersions)

	return versions
}
